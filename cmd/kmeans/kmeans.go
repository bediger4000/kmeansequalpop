package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"kmeansequalpop/centroids"
	"kmeansequalpop/data"
)

func main() {
	points, err := data.ReadPoints(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# %d points in input\n", len(points))

	k, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# %d clusters\n", k)

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	centers := centroids.KmeansPPCentroids(points, k)

	fmt.Printf("# Found %d initial k-means++ centroids\n", len(centers))

	for i := range centers {
		fmt.Printf("# Centroid %d at (%f,%f)\n", i, centers[i].X, centers[i].Y)
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
			totalPopulation += cluster[i].Pop
			fmt.Printf("%f	%f	%d\n", cluster[j].X, cluster[j].Y, i)
		}
		fmt.Printf("# cluster %d, %d items, %.0f population\n", i, len(cluster), population)
	}
	fmt.Printf("# %d total points, %f total population\n", totalPoints, totalPopulation)

	// 1. Compute current cluster means
	var clusterCenters []*centroids.Coords
	for _, cluster := range clusters {
		center := centroids.CalculateCentroid(cluster)
		clusterCenters = append(clusterCenters, center)
	}
	for _, center := range clusterCenters {
		fmt.Printf("%f %f NC\n", center.X, center.Y)
	}
	// 2. For each object, compute the distances to the cluster means
	distances := centroids.ClusterCenterDistances(points, clusterCenters)
	fmt.Printf("# %d distances to cluster means\n", len(distances))

	// 3. Sort elements based on the delta of the current assignment and the best
	//    possible alternate assignment.
	distances = centroids.SortByAssignment(distances)
	fmt.Printf("# %d sorted distances to cluster means\n", len(distances))

	/* 4. For each element by priority:
	   1. For each other cluster, by element gain, unless already moved:
		  1. If there is an element wanting to leave the other cluster and this
		     swap yields and improvement, swap the two elements
		  2. If the element can be moved without violating size constraints,
		     move it
	   2. If the element was not changed, add to outgoing transfer list.
	*/

	/*
		for i := range distances {
			dist := distances[i]
			pt := dist.Point
		}
	*/

	// 5. If no more transfers were done (or max iteration threshold was reached), terminate
}
