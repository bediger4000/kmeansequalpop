package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

	centroids := centroids.KmeansPPCentroids(points, k)

	fmt.Printf("# Found %d initial k-means++ centroids\n", len(centroids))
}
