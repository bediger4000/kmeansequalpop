package centroids

import (
	"kmeansequalpop/data"
	"math"
	"sort"
)

// CalculateCentroid conjures the (X,Y) position of the centroid
// of a cluster of *data.Point. Population-weighted.
func CalculateCentroid(points []*data.Point) *Coords {

	sumXmoments := 0.0
	sumYmoments := 0.0
	sumPopulation := 0.0

	for i := range points {
		sumXmoments += points[i].Xmoment
		sumYmoments += points[i].Ymoment
		sumPopulation += points[i].Pop
		points[i].Assigned = false
	}

	// X coord of centroid is sum of moments around Y-axis/population
	// Think about it, coordinate with pre-calculated moments in func ReadPoints
	return &Coords{X: sumYmoments / sumPopulation, Y: sumXmoments / sumPopulation}
}

// SortByAssignment sorts elements based on the delta of the current
// assignment and the best possible alternate assignment.  Looks like
// this delta could be negative, if the best possible alternate
// assignment is farther than the point's current centroid.
func SortByAssignment(distances []*PointDistance) []*PointDistance {
	// Find delta of current assignment and the best
	// possible alternate assignment.
	for i := range distances {
		pt := distances[i].Point
		dx := distances[i].CentroidX - pt.X
		dy := distances[i].CentroidY - pt.Y
		currentDist := math.Sqrt(dx*dx + dy*dy)
		distances[i].DistWeight = pt.Pop * (currentDist - distances[i].MinDist)
	}
	sort.Sort(opSlice(distances))

	return distances
}

// ClusterCenterDistances - for each object, compute the distances to the cluster means
func ClusterCenterDistances(points []*data.Point, centers []*Coords) []*PointDistance {
	var ccdistances []*PointDistance

	for i := range points {
		dist := &PointDistance{Point: points[i], MinDist: math.MaxFloat32}

		for _, center := range centers {
			dx := points[i].X - center.X
			dy := points[i].Y - center.Y

			distance := math.Sqrt(dx*dx + dy*dy)

			if distance < dist.MinDist {
				dist.MinDist = distance
				dist.CentroidX = center.X
				dist.CentroidY = center.Y
			}

			// dist.Distances = append(dist.Distances, distance)
		}

		ccdistances = append(ccdistances, dist)
	}

	return ccdistances
}
