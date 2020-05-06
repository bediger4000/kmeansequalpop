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

	distances := orderPointsByDistance(points, centers)

	clusters := make([][]*data.Point, k)
	var fullClusters [][]*data.Point
	// first index of clusters, clusters[i], connects to centers[i]
	clusterPopulations := make([]float64, k)

	// Assign points to their preferred cluster until this cluster
	// is full, then resort remaining objects, without taking the
	// full cluster into account anymore.

	for len(distances) > 0 {

		for i := range distances {

			minDistIdx := distances[i].minDistIndex()
			pt := distances[i].Point

			clusters[minDistIdx] = append(clusters[minDistIdx], pt)
			clusterPopulations[minDistIdx] += pt.Pop
			pt.Assigned = true

			if clusterPopulations[minDistIdx] >= desiredClusterPopulation {
				fullClusters = append(fullClusters, clusters[minDistIdx])

				clusters = append(clusters[:minDistIdx], clusters[minDistIdx+1:]...)
				centers = append(centers[:minDistIdx], centers[minDistIdx+1:]...)
				clusterPopulations = append(clusterPopulations[:minDistIdx], clusterPopulations[minDistIdx+1:]...)

				break
			}
		}

		distances = orderPointsByDistance(points, centers)
	}

	// One cluster will be left unfilled, probably?
	fullClusters = append(fullClusters, clusters...)

	return fullClusters
}

type PointDistance struct {
	DistDiff  float64
	Point     *data.Point
	Distances []float64
}

type opSlice []*PointDistance

func (ps opSlice) Len() int { return len(ps) }
func (ps opSlice) Less(i, j int) bool {
	if ps[i].DistDiff > ps[j].DistDiff {
		return true
	}
	return false
}
func (ps opSlice) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

// orderPointsByDistance orders Points by the distance to their nearest
// cluster minus distance to the farthest cluster (= biggest benefit of
// best over worst assignment
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
			dist := dx*dx + dy*dy
			pd.Distances[centerIndex] = dist
			if dist > maxDist {
				maxDist = dist
			}
			if dist < minDist {
				minDist = dist
			}
		}

		pd.DistDiff = maxDist - minDist

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
	for j, dist := range pd.Distances {
		if dist < minDist {
			minDist = dist
			minDistIdx = j
		}
	}
	return minDistIdx
}
