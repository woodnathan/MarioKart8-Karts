package main

import (
	"./mk8data"
	"fmt"
)

type Combination struct {
	mk8data.Record
	Character mk8data.Character
	Body mk8data.Body
	Tire mk8data.Tire
	Glider mk8data.Glider
}

func (self *Combination) calculateSpeed() {
	self.Speed.Value = self.Character.Speed.Value + self.Body.Speed.Value + self.Tire.Speed.Value + self.Glider.Speed.Value
	self.Speed.Water = self.Character.Speed.Water + self.Body.Speed.Water + self.Tire.Speed.Water + self.Glider.Speed.Water
	self.Speed.Air = self.Character.Speed.Air + self.Body.Speed.Air + self.Tire.Speed.Air + self.Glider.Speed.Air
	self.Speed.Ground = self.Character.Speed.Ground + self.Body.Speed.Ground + self.Tire.Speed.Ground + self.Glider.Speed.Ground
}
func (self *Combination) calculateTotal() {
	self.calculateSpeed()
	self.Acceleration = self.Character.Acceleration + self.Body.Acceleration + self.Tire.Acceleration + self.Glider.Acceleration
}
func (self *Combination) name() string {
	return fmt.Sprintf("%s, %s, %s, %s", self.Character.Name, self.Body.Name, self.Tire.Name, self.Glider.Name)
}

func createCombinations(dir string) (error) {
	characters, err := mk8data.ReadCharacters(dir + "characters.json")
	bodies, err := mk8data.ReadBodies(dir + "bodies.json")
	tires, err := mk8data.ReadTires(dir + "tires.json")
	gliders, err := mk8data.ReadGliders(dir + "gliders.json")
	if err != nil {
		return err
	}
	
	perm_count := len(characters) * len(bodies) * len(tires) * len(gliders)
	
	combinations := make([]Combination, perm_count)
	i := 0
	for _, character := range(characters) {
		for _, body := range(bodies) {
			for _, tire := range(tires) {
				for _, glider := range(gliders) {
					combination := &combinations[i]
					
					combination.Character = character
					combination.Body = body
					combination.Tire = tire
					combination.Glider = glider
					
					combination.calculateTotal()
					
					combinations[i] = *combination
					i++
				}
			}
		}
	}
	
	var c *Combination
	for i, combination := range(combinations) {
		if c == nil || combination.Acceleration > c.Acceleration {
			c = &combinations[i]
		}
	}
	fmt.Printf("Acceleration:\t%.2f = %s\n", c.Acceleration, c.name())
	
	for i, combination := range(combinations) {
		if c == nil || combination.Speed.Value > c.Speed.Value {
			c = &combinations[i]
		}
	}
	fmt.Printf("Speed:\t\t%.2f = %s\n", c.Speed.Value, c.name())
	
	for i, combination := range(combinations) {
		if c == nil || combination.Speed.Water > c.Speed.Water {
			c = &combinations[i]
		}
	}
	fmt.Printf("Speed (Water):\t%.2f = %s\n", c.Speed.Water, c.name())
	
	for i, combination := range(combinations) {
		if c == nil || combination.Speed.Air > c.Speed.Air {
			c = &combinations[i]
		}
	}
	fmt.Printf("Speed (Air):\t%.2f = %s\n", c.Speed.Air, c.name())
	
	for i, combination := range(combinations) {
		if c == nil || combination.Speed.Ground > c.Speed.Ground {
			c = &combinations[i]
		}
	}
	fmt.Printf("Speed (Ground):\t%.2f = %s\n", c.Speed.Ground, c.name())
	
	return nil
}

func main() {
	err := createCombinations("MarioKart8-Stats/json/")
	if err != nil {
		fmt.Println(err)
	}
}