package main

import (
	"encoding/json"
	"os"
)

// Airport is the struct for the airports
type Airport struct {
	Name    string `json:"nome"`
	Airport string `json:"aeroporto"`
	City    string `json:"cidade"`
}

// Airports array of airport
type Airports []Airport

// LoadAirportsFromFile with read a json file of airport and return an array
func LoadAirportsFromFile(filename string) (Airports, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	airports := make(Airports, 0)

	err = json.NewDecoder(f).Decode(&airports)
	return airports, err
}
