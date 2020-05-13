package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"kmeansequalpop/centroids"
	"kmeansequalpop/data"
)

func main() {
	points, population, err := data.ReadPoints(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# %d points in input, %d total populatoin\n", len(points), population)

	k, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# %d clusters\n", k)

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	centers := centroids.KmeansPPCentroids(points, k)

	fmt.Printf("# Found %d initial k-means++ centroids\n", len(centers))

	for i := range centers {
		fmt.Printf("# Centroid %d at (%f,%f)\n", centers[i].ClusterIdx, centers[i].X, centers[i].Y)
		fmt.Printf("%f %f c\n", centers[i].X, centers[i].Y)
	}

	clusters := centroids.InitialAssign(points, centers, k)

	fmt.Printf("# Assigned points to %d clusters\n", len(clusters))
	totalPopulation := 0.0
	totalPoints := 0
	for i, cluster := range clusters {
		population := 0.0
		totalPoints += len(cluster)
		for j := range cluster {
			population += cluster[j].Pop
			//fmt.Printf("%f	%f	%d\n", cluster[j].X, cluster[j].Y, i)
		}
		fmt.Printf("# cluster %d, %d items, %.0f population\n", i, len(cluster), population)
		totalPopulation += population
	}
	fmt.Printf("# %d total points, %.0f total population\n", totalPoints, totalPopulation)

	// Loop Top
	for iteration := 0; iteration < 20; iteration++ {
		// 1. Compute current cluster means
		var clusterCenters []*centroids.Coords
		for i, cluster := range clusters {
			center := centroids.CalculateCentroid(cluster)
			center.ClusterIdx = i
			clusterCenters = append(clusterCenters, center)
		}
		for idx, center := range clusterCenters {
			fmt.Printf("# iteration %d %d %f %f NC\n", iteration, idx, center.X, center.Y)
		}
		// 2. For each object, compute the distances to the cluster means
		distances := centroids.ClusterCenterDistances(points, clusterCenters)
		fmt.Printf("# iteration %d %d distances to cluster means\n", iteration, len(distances))

		// 3. Sort elements based on the delta of the current assignment and the best
		//    possible alternate assignment.
		distances = centroids.SortByAssignment(distances)
		fmt.Printf("# iteration %d %d sorted distances to cluster means\n", iteration, len(distances))

		/* 4. For each element by priority:
		   1. For each other cluster, by element gain, unless already moved:
			  1. If there is an element wanting to leave the other cluster and this
			     swap yields and improvement, swap the two elements
			  2. If the element can be moved without violating size constraints,
			     move it
		   2. If the element was not changed, add to outgoing transfer list.
		*/

		// track populations of clusters dynamically
		clusterPopulations := centroids.CountPopulation(distances, k)
		minClusterPop := math.MaxFloat32
		maxClusterPop := 0.0
		totalPopulation = 0.0
		for i, pop := range clusterPopulations {
			totalPopulation += pop
			fmt.Printf("#A iteration %d cluster %d, population %.0f\n", iteration, i, pop)
			if pop > maxClusterPop {
				maxClusterPop = pop
			}
			if pop < minClusterPop {
				minClusterPop = pop
			}
		}
		fmt.Printf("#A iteration %d total population %.0f\n", iteration, totalPopulation)

		var transferCount int
		leavers := centroids.WantToLeave(distances, k)

		transferCount = 0

		// for each element by priority
		for i := range distances {
			dist := distances[i]
			pt := dist.Point

		CLUSTERLOOP:
			for j := range clusters {
				// for every other cluster
				if j == pt.CentroidIdx {
					continue
				}
				for k := range leavers[j] {
					// If there is an element wanting to leave the other cluster
					// and this swap yields an improvement, swap the two
					// elements
					leaver := leavers[j][k]

					if leaver.Point.X == pt.X && leaver.Point.Y == pt.Y &&
						leaver.Point.Pop == pt.Pop {
						// leaver and pt are the same Point
						continue
					}

					if pt.CentroidIdx == leaver.Point.CentroidIdx {
						// same cluster, don't bother
						continue
					}

					d1 := leaver.Distances[leaver.Point.CentroidIdx]
					d2 := leaver.Distances[pt.CentroidIdx]

					if d2 < d1 {
						// if one population goes over the limit,
						// put this point on leavers list for its cluster,
						// then continue this loop. Don't swap items.
						ptPop := clusterPopulations[pt.CentroidIdx] - pt.Pop + leaver.Point.Pop
						lvPop := clusterPopulations[leaver.Point.CentroidIdx] - leaver.Point.Pop + pt.Pop
						// fmt.Printf("# Pt pop %.0f, leave pop %.0f\n", ptPop, lvPop)
						if ptPop < minClusterPop || lvPop < minClusterPop || ptPop > maxClusterPop || lvPop > maxClusterPop {
							// If the element was not changed, add to outgoing transfer list.
							leavers[j] = append(leavers[j], dist)
							break CLUSTERLOOP
						}

						// Adjust clusters' populations
						clusterPopulations[leaver.Point.CentroidIdx] -= (leaver.Point.Pop - pt.Pop)
						clusterPopulations[pt.CentroidIdx] -= (pt.Pop - leaver.Point.Pop)

						// Swap centroids
						pt.CentroidIdx, leaver.Point.CentroidIdx = leaver.Point.CentroidIdx, pt.CentroidIdx

						// remove from leavers list
						leavers[j] = append(leavers[j][:k], leavers[j][k+1:]...)
						transferCount++
						break // out of range over leavers
					}
				}
			}
		}

		fmt.Printf("# iteration %d transfers %d\n", iteration, transferCount)

		// Need to rework clusters - some points changed cluster
		clusters = centroids.AssignToCluster(points, k)

		totalPopulation = 0.0
		for i, cluster := range clusters {
			population := 0.0
			for j := range cluster {
				population += cluster[j].Pop
				fmt.Printf("iteration %d %f	%f	%d\n", iteration, cluster[j].X, cluster[j].Y, i)
			}
			totalPopulation += population
			fmt.Printf("#X iteration %d cluster %d, %d items, %.0f population\n", iteration, i, len(cluster), population)
		}
		fmt.Printf("# iteration %d Total population %.0f\n", iteration, totalPopulation)

		// 5. If no more transfers were done (or max iteration threshold was reached), terminate
	}
}
