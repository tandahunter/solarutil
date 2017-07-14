//Package solarutil contains utility functions for working with spatial bodies
package solarutil

//PlanetArray stores an array of planets
type PlanetArray []Planet

//NewPlanetArray returns a pointer to a new planet array
func NewPlanetArray() *PlanetArray {
	return &PlanetArray{}
}

func (p PlanetArray) printNames() {
	for _, planet := range p {
		planet.printName()
	}
}

//GetPlanets returns a pointer to all planets where IsStar = false
func (p PlanetArray) GetPlanets() *PlanetArray {
	return p.getFilteredPlanets(false)
}

//GetStars returns a pointer to all stars where IsStar = true
func (p PlanetArray) GetStars() *PlanetArray {
	return p.getFilteredPlanets(true)
}

//FirstOrDefault returns a pointer to the first or a new planet in the array
func (p PlanetArray) FirstOrDefault() *Planet {
	if len(p) > 0 {
		return &p[0]
	}

	return &Planet{}
}

func (p PlanetArray) getFilteredPlanets(isStar bool) *PlanetArray {
	toReturn := PlanetArray{}

	for _, planet := range p {
		if planet.IsStar == isStar {
			toReturn = append(toReturn, planet)
		}
	}

	return &toReturn
}
