//Package solarutil contains utility functions for working with spatial bodies
package solarutil

//PlanetArray stores an array of planets
type PlanetArray []Planet

//NewPlanetArray returns a pointer to a new planet array
func NewPlanetArray() *PlanetArray {
	return &PlanetArray{}
}

//PrintNames prints the name of all planets in the array
func (p PlanetArray) PrintNames() {
	for _, planet := range p {
		planet.PrintName()
	}
}

//GetPlanets returns a pointer to all planets where IsStar = false
func (p PlanetArray) GetPlanets() *PlanetArray {
	return p.getFilteredPlanets(false, "")
}

//GetStars returns a pointer to all stars where IsStar = true
func (p PlanetArray) GetStars() *PlanetArray {
	return p.getFilteredPlanets(true, "")
}

//GetPlanetsByName returns a pointer to all stars where Name = name
func (p PlanetArray) GetPlanetsByName(name string) *PlanetArray {
	return p.getFilteredPlanets(false, name)
}

//FirstOrDefault returns a pointer to the first or a new planet in the array
func (p PlanetArray) FirstOrDefault() *Planet {
	if len(p) > 0 {
		return &p[0]
	}

	return &Planet{}
}

func (p PlanetArray) getFilteredPlanets(isStar bool, name string) *PlanetArray {
	toReturn := PlanetArray{}

	for _, planet := range p {
		if planet.IsStar == isStar && (name == "" || planet.Name == name) {
			toReturn = append(toReturn, planet)
		}
	}

	return &toReturn
}
