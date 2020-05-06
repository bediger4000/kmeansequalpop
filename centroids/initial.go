package centroids

import (
	"kmeansequalpop/data"
	"math/rand"
)

// Coords holds X, Y cartesian coordinates of a cluster's centroid
type Coords struct {
	X float64
	Y float64
}

// KmeansPPCentroids chooses k elements of points as initial centroids
// via the k-means++ algorithm
func KmeansPPCentroids(points []*data.Point, k int) []Coords {
	var centroids []Coords

	// 1. Choose one center uniformly at random among the data points.
	pt := points[rand.Intn(len(points)-1)]
	centroids = append(centroids, Coords{X: pt.X, Y: pt.Y})

	for i := 1; i < k; i++ {
		// 2. For each data point x, compute D(x), the distance between x and
		// the nearest center that has already been chosen.
		distances := computeD2(points, centroids)

		// 3. Choose one new data point at random as a new center, using a
		// weighted probability distribution where a point x is chosen with
		// probability proportional to D(x)^2.
		idx := chooseByD2(distances)
		pt = points[idx]
		centroids = append(centroids, Coords{X: pt.X, Y: pt.Y})

	}

	return centroids
}
func computeD2(points []*data.Point, centroids []Coords) []float64 {

	distances := make([]float64, len(points))

	for i := range points {
		pt := points[i]

		dx := centroids[0].X - pt.X
		dy := centroids[0].Y - pt.Y

		minD2 := dx*dx + dy*dy

		for j := 1; j < len(centroids); j++ {
			dx = centroids[j].X - pt.X
			dy = centroids[j].Y - pt.Y
			D2 := dx*dx + dy*dy
			if D2 < minD2 {
				minD2 = D2
			}
		}

		distances[i] = minD2
	}

	return distances
}

// chooseByD2 returns an int, the index of a data.Point
// whose dist^2 to the nearest existing centroid is in
// slice distances. The choice should be random, but weighted
// by the size of dist^2
func chooseByD2(distances []float64) int {
	intervals := make([]float64, len(distances))

	sum := 0.0
	for i := 0; i < len(distances); i++ {
		sum += distances[i]
		intervals[i] = sum
	}
	// intervals contains sorted values: they just increase
	// as index in intervals gets bigger

	// inInterval should appear uniformly randomly throughout the range of
	// dist^2 values. Find the smallest value of intervals less than the value
	// of inInterval.
	inInterval := rand.Float64() * intervals[len(intervals)-1]

	for i := range intervals {
		if inInterval < intervals[i] {
			return i
		}
	}

	return 0
}
