package data

// Point an x,y cartesian point
type Point struct {
	Pop         float64
	X           float64
	Y           float64
	Assigned    bool
	Xmoment     float64
	Ymoment     float64
	CentroidIdx int // cluster membership
}
