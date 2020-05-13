package centroids

// WantToLeave creates lists of *PointDistance, one per
// cluster, of items that might "want" to leave the cluster.
func WantToLeave(distances []*PointDistance, k int) [][]*PointDistance {
	leavers := make([][]*PointDistance, k)

	for i := range distances {
		centroidIdx := distances[i].Point.CentroidIdx
		currentDist := distances[i].Distances[centroidIdx]

		for _, dist := range distances[i].Distances {
			if dist < currentDist {
				leavers[centroidIdx] = append(leavers[centroidIdx], distances[i])
				break
			}
		}
	}

	return leavers
}
