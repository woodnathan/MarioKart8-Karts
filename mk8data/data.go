package mk8data

import (
	"os"
	"encoding/json"
)

type MultipartElement struct {
	Value float64 `json:"value"`
	Water float64 `json:"water"`
	Air float64 `json:"air"`
	Ground float64 `json:"ground"`
}

type Speed struct {
	MultipartElement
}
type Handling struct {
	MultipartElement
}

type Record struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Speed Speed `json:"speed"`
	Acceleration float64 `json:"acceleration"`
	Weight float64 `json:"weight"`
	Handling Handling `json:"handling"`
	Traction float64 `json:"traction"`
	MiniTurbo float64 `json:"mini_turbo"`
}

type Character struct {
	Record
}
type Body struct {
	Record
}
type Tire struct {
	Record
}
type Glider struct {
	Record
}

func ReadCharacters(path string) ([]Character, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var characters []Character
	err = json.NewDecoder(file).Decode(&characters)
	file.Close()
	
	return characters, err
}

func ReadBodies(path string) ([]Body, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var bodies []Body
	err = json.NewDecoder(file).Decode(&bodies)
	file.Close()
	
	return bodies, err
}

func ReadTires(path string) ([]Tire, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var tires []Tire
	err = json.NewDecoder(file).Decode(&tires)
	file.Close()
	
	return tires, err
}

func ReadGliders(path string) ([]Glider, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var gliders []Glider
	err = json.NewDecoder(file).Decode(&gliders)
	file.Close()
	
	return gliders, err
}
