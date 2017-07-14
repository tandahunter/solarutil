//Package solarutil contains utility functions for working with spatial bodies
package solarutil

//Vector stores coordinates in 3d space
type Vector struct {
	X float64
	Y float64
	Z float64
}

//NewVector returns a vector struct
func NewVector(x, y, z float64) *Vector {
	vector := Vector{}
	vector.X = x
	vector.Y = y
	vector.Z = z

	return &vector
}
