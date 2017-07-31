//Package solarutil contains utility functions for working with spatial bodies
package solarutil

//PlanetArray stores an array of pointers to planets
type PlanetArray []*Planet

//NewPlanetArray returns a pointer to a new planet array
func NewPlanetArray() *PlanetArray {
	return &PlanetArray{}
}

//GetPlanetByID returns a pointer to all stars where Name = name
func (p *PlanetArray) GetPlanetByID(id int) *Planet {

	for _, planet := range *p {
		if planet.ID == id {
			return planet
		}
	}

	return nil
}

//AddNewPlanet creates a new planet and adds a pointer to the array
func (p *PlanetArray) AddNewPlanet(id int, name string, mass, distance, speed float64) *Planet {
	planet := NewPlanet(id, name, mass, distance, speed)
	*p = append(*p, planet)
	return planet
}

//AddNewTexturedPlanet creates a new textured planet and adds a pointer to the array
func (p *PlanetArray) AddNewTexturedPlanet(id int, name string, mass, distance, speed float64, texture string) *Planet {
	planet := NewPlanet(id, name, mass, distance, speed)
	planet.Texture = texture
	*p = append(*p, planet)
	return planet
}
