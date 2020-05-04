package centroids

import "kmeansequalpop/data"

type Coords struct {
	X float64
	Y float64
}

// KmeansPPCentroids chooses k elements of points as initial centroids
// via the k-means++ algorithm
func KmeansPPCentroids(points []*data.Point, k int) []Coords {
	return []Coords{}
}
