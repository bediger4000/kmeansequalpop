package main

import (
	"fmt"
	"log"
	"os"

	"kmeansequalpop/data"
)

func main() {
	points, err := data.ReadPoints(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# %d points in input\n", len(points))
	for i := range points {
		fmt.Printf("%.0f %f %f\n", points[i].Pop, points[i].X, points[i].Y)
	}
}
