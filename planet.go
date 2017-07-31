//Package solarutil contains utility functions for working with spatial bodies
package solarutil

import (
	"math"
)

//Planet stores information about a planetary body
type Planet struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Mass     float64 `json:"mass"`
	Vector   *Vector `json:"vector"`
	Velocity *Vector `json:"velocity"`
	Texture  string  `json:"texture"`
}

//NewPlanet returns a pointer to new planet
func NewPlanet(id int, name string, mass, distance, speed float64) *Planet {
	planet := Planet{}
	planet.ID = id
	planet.Name = name
	planet.Mass = mass
	planet.Vector = NewVector(distance, 0, 0)
	planet.Velocity = NewVector(0, 0, speed)

	return &planet
}

//DistanceTo returns the distance between this and the specified star.
func (p *Planet) DistanceTo(s *Star) float64 {
	vectorA := *p.Vector
	vectorB := *s.Vector
	pow := float64(2)

	x := math.Pow(vectorA.X-vectorB.X, pow)
	y := math.Pow(vectorA.Y-vectorB.Y, pow)
	z := math.Pow(vectorA.Z-vectorB.Z, pow)

	return math.Sqrt(x + y + z)
}
