package centroids

func CountPopulation(distances []*PointDistance, k int) []float64 {
	populations := make([]float64, k)
	for i := range distances {
		populations[distances[i].Point.CentroidIdx] += distances[i].Point.Pop
	}
	return populations
}
