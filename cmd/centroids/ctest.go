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

	fmt.Printf("# %d cluster\n", k)

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
		for i := range cluster {
			population += cluster[i].Pop
			totalPopulation += cluster[i].Pop
		}
		fmt.Printf("# cluster %d, %d items, %.0f population\n", i, len(cluster), population)
	}
	fmt.Printf("# %d total points, %f total population\n", totalPoints, totalPopulation)

	for i := range clusters {
		for j := range clusters[i] {
			pt := clusters[i][j]
			fmt.Printf("%f %f %.0f %d\n", pt.X, pt.Y, pt.Pop, i)
		}
	}
}
