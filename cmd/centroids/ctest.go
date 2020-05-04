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

	centroids := centroids.KmeansPPCentroids(points, k)

	fmt.Printf("# Found %d initial k-means++ centroids\n", len(centroids))

	for i := range centroids {
		fmt.Printf("# Centroid %d at (%f,%f)\n", i, centroids[i].X, centroids[i].Y)
	}
}
