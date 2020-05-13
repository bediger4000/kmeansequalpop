package centroids

import "kmeansequalpop/data"

// PointDistance helps calculate which initial centroid in which
// to place a data.Point.
type PointDistance struct {
	DistWeight float64
	Point      *data.Point
	Distances  []float64
	MinDist    float64
	CentroidX  float64
	CentroidY  float64
}

type opSlice []*PointDistance

// Coords holds X, Y cartesian coordinates of a cluster's centroid
type Coords struct {
	X          float64
	Y          float64
	ClusterIdx int
}

func (ps opSlice) Len() int { return len(ps) }
func (ps opSlice) Less(i, j int) bool {
	if ps[i].DistWeight > ps[j].DistWeight {
		return true
	}
	return false
}
func (ps opSlice) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }
