//Package solarutil contains utility functions for working with spatial bodies
package solarutil

import (
	"fmt"
	"math"
)

//Planet stores information about a planetary body
type Planet struct {
	Name   string
	Mass   float64
	Vector *Vector
	IsStar bool
}

//NewPlanet returns a pointer to new planet
func NewPlanet(name string, mass float64, x, y, z float64) *Planet {
	planet := Planet{}
	planet.Name = name
	planet.Mass = mass
	planet.Vector = NewVector(x, y, z)
	planet.IsStar = false

	return &planet
}

//NewStar returns pointer to new star
func NewStar(name string, mass float64, x, y, z float64) *Planet {
	planet := NewPlanet(name, mass, x, y, z)
	planet.IsStar = true

	return planet
}

func (p *Planet) printName() {
	fmt.Print(p.Name)
	fmt.Println()
}

//DistanceTo returns the distance between this and the specified body.
func (p *Planet) DistanceTo(p2 *Planet) float64 {
	vectorA := *p.Vector
	vectorB := *p2.Vector
	pow := float64(2)

	x := math.Pow(vectorA.X-vectorB.X, pow)
	y := math.Pow(vectorA.Y-vectorB.Y, pow)
	z := math.Pow(vectorA.Z-vectorB.Z, pow)

	return math.Sqrt(x + y + z)
}
