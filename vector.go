//Package solarutil contains utility functions for working with spatial bodies
package solarutil

import "math"

//Vector stores coordinates in 3d space
type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

//NewVector returns a vector struct
func NewVector(x, y, z float64) *Vector {
	vector := Vector{}
	vector.X = x
	vector.Y = y
	vector.Z = z

	return &vector
}

//SubVectors assigns the difference between two vectors to this
//vector's coordinates
func (v *Vector) SubVectors(a, b *Vector) {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
}

//GetLength calculates the length of this vector....???
func (v *Vector) GetLength() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

//SetLength sets the length of this vector... this doesn't make sense.
func (v *Vector) SetLength(a float64) {
	b := v.GetLength()

	if b != 0 && a != b {
		v.MultiplyScalar(a / b)
	}
}

//MultiplyScalar applies a multiplication transformation to x,y,z
func (v *Vector) MultiplyScalar(a float64) {
	v.X *= a
	v.Y *= a
	v.Z *= a
}

//Add adds a vector to this.  (x = x + a.x etc)
func (v *Vector) Add(a *Vector) {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
}
