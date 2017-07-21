//Package solarutil contains utility functions for working with spatial bodies
package solarutil

import (
	"github.com/twinj/uuid"
)

//Star stores information about a star
type Star struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Mass   float64 `json:"mass"`
	Vector *Vector `json:"vector"`
}

//NewStar returns a pointer to new planet
func NewStar(name string, mass float64, x, y, z float64) *Star {
	star := Star{}
	star.ID = uuid.NewV4().String()
	star.Name = name
	star.Mass = mass
	star.Vector = NewVector(x, y, z)

	return &star
}
