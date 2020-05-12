package centroids

import (
	"kmeansequalpop/data"
	"math"
	"sort"
)

// InitialAssign puts data.Point instances into their initial
// cluster assignment.
func InitialAssign(points []*data.Point, centers []Coords, k int) [][]*data.Point {
	totalPopulace := 0.0
	for i := range points {
		totalPopulace += points[i].Pop
	}
	desiredClusterPopulation := totalPopulace / float64(k)

	weightedDistances := orderPointsByDistance(points, centers)

	clusters := make([][]*data.Point, k)
	var fullClusters [][]*data.Point
	// first index of clusters, clusters[i], connects to centers[i]
	clusterPopulations := make([]float64, k)

	// Assign points to their preferred cluster until this cluster
	// is full, then resort remaining objects, without taking the
	// full cluster into account anymore.

	for len(weightedDistances) > 0 {

		for i := range weightedDistances {

			// Unlike weightedDistances, minDistIndex() does give
			// back index of minimum distance to a centroid, because
			// we're working on a single point.
			minDistIdx := weightedDistances[i].minDistIndex()
			pt := weightedDistances[i].Point

			clusters[minDistIdx] = append(clusters[minDistIdx], pt)
			clusterPopulations[minDistIdx] += pt.Pop
			pt.Assigned = true
			pt.CentroidX = centers[minDistIdx].X
			pt.CentroidY = centers[minDistIdx].Y

			if clusterPopulations[minDistIdx] >= desiredClusterPopulation {
				fullClusters = append(fullClusters, clusters[minDistIdx])

				clusters = append(clusters[:minDistIdx], clusters[minDistIdx+1:]...)
				centers = append(centers[:minDistIdx], centers[minDistIdx+1:]...)
				clusterPopulations = append(clusterPopulations[:minDistIdx], clusterPopulations[minDistIdx+1:]...)

				break
			}
		}

		// re-sort remaining objects, without taking the
		// full cluster into account.
		weightedDistances = orderPointsByDistance(points, centers)
	}

	// One cluster will be left unfilled, probably?
	fullClusters = append(fullClusters, clusters...)

	return fullClusters
}

// orderPointsByDistance orders Points by the distance to their nearest
// cluster minus distance to the farthest cluster (= biggest benefit of
// best over worst assignment.
func orderPointsByDistance(points []*data.Point, centers []Coords) []*PointDistance {
	var orderedPoints []*PointDistance

	for i := range points {
		pt := points[i]
		if pt.Assigned {
			continue
		}

		pd := &PointDistance{Point: pt}
		pd.Distances = make([]float64, len(centers))

		maxDist := -1.0
		minDist := math.MaxFloat32
		for centerIndex, center := range centers {
			dx := center.X - pt.X
			dy := center.Y - pt.Y
			dist := math.Sqrt(dx*dx + dy*dy)
			pd.Distances[centerIndex] = dist
			if dist > maxDist {
				maxDist = dist
			}
			if dist < minDist {
				minDist = dist
			}
		}

		pd.DistWeight = (maxDist - minDist)

		orderedPoints = append(orderedPoints, pd)
	}

	sort.Sort(opSlice(orderedPoints))

	return orderedPoints
}

// minDistIndex calculates the index at which the referenced
// point has a minimum distance to a cluster centroid.
func (pd *PointDistance) minDistIndex() int {
	minDist := math.MaxFloat32
	var minDistIdx int
	for i, dist := range pd.Distances {
		if dist < minDist {
			minDist = dist
			minDistIdx = i
		}
	}
	return minDistIdx
}
