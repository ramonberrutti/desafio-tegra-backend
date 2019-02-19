package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Flight represent the struct for a flight
type Flight struct {
	Flight        string    `json:"voo"`
	Origin        string    `json:"origem"`
	Destination   string    `json:"destino"`
	DataDeparture string    `json:"data_saida"`
	TimeDeparture string    `json:"saida"`
	TimeArrival   string    `json:"chegada"`
	Price         float32   `json:"valor"`
	Company       string    `json:"operadora"`
	Departure     time.Time `json:"departure"`
	Arrival       time.Time `json:"arrival"`
}

// Flights array of Flights
type Flights []Flight

var errInvalidFormat = errors.New("Invalid file format")

// LoadFlightsFromFile will load flights from an specific company
func LoadFlightsFromFile(filename, company string) (Flights, error) {
	// Check Extension
	fileExt := filepath.Ext(filename)

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if fileExt == ".json" {
		return loadJSON(f, company)
	} else if fileExt == ".csv" {
		return loadCSV(f, company)
	}

	return nil, errInvalidFormat
}

func loadCSV(f *os.File, company string) (Flights, error) {
	r := csv.NewReader(f)

	// read header
	if _, err := r.Read(); err != nil {
		return nil, err
	}

	flights := make(Flights, 0)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		// strconv.FormatFloat(line[6], 'f', 2, 32)
		price, err := strconv.ParseFloat(line[6], 32)
		if err != nil {
			continue
		}

		tDepature, _ := convertDateToTime(line[3], line[4])
		tArrival, _ := convertDateToTime(line[3], line[5])

		flights = append(flights, Flight{
			Company:       company,
			Flight:        line[0],
			Origin:        line[1],
			Destination:   line[2],
			DataDeparture: line[3],
			TimeDeparture: line[4],
			TimeArrival:   line[5],
			Price:         float32(price),
			Departure:     tDepature,
			Arrival:       tArrival,
		})

	}

	return flights, nil
}

func loadJSON(f *os.File, company string) (Flights, error) {
	flights := make(Flights, 0)

	err := json.NewDecoder(f).Decode(&flights)
	if err != nil {
		return nil, err
	}

	// Set the Company, this don't is very performant :cry:
	for i := range flights {
		flights[i].Company = company

		flights[i].Departure, err = convertDateToTime(flights[i].DataDeparture, flights[i].TimeDeparture)
		flights[i].Arrival, _ = convertDateToTime(flights[i].DataDeparture, flights[i].TimeArrival)
	}

	return flights, nil
}
