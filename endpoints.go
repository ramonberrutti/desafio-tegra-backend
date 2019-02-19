package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"
)

type jsonPatch struct {
	Flight      string    `json:"voo"`
	Origin      string    `json:"origem"`
	Destination string    `json:"destino"`
	Departure   time.Time `json:"saida"`
	Arrival     time.Time `json:"chegada"`
	Company     string    `json:"operadora"`
	Price       float32   `json:"preco"`
}

type jsonFullFlight struct {
	Origin      string      `json:"origem"`
	Destination string      `json:"destino"`
	Departure   time.Time   `json:"saida"`
	Arrival     time.Time   `json:"chegada"`
	Patches     []jsonPatch `json:"trechos"`
}

func airportList(airports Airports, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airports)
}

func searchFlights(graph *Graph, w http.ResponseWriter, r *http.Request) {

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	date := r.URL.Query().Get("date")

	if from == "" || to == "" || date == "" {
		http.Error(w, "Invalid Params", http.StatusBadRequest)
		return
	}

	tDate, err := convertDateToTime(date, "00:00")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	routes := graph.FoundRoute(from, to, func(nameFrom string, edge Edge) bool {
		return compareDate(edge.Departure, tDate)
	}, func(nameFrom string, edgeFrom, edge Edge) bool {
		diff := edge.Departure.Sub(edgeFrom.Arrival)

		return diff >= 0 && diff <= time.Hour*12
	})

	// log.Println(routes)
	flights := make([]jsonFullFlight, 0)
	for i := range routes {
		patches := make([]jsonPatch, 0)

		for j := range routes[i] {
			patches = append(patches, jsonPatch{
				Company:     routes[i][j].Company,
				Flight:      routes[i][j].Flight,
				Origin:      routes[i][j].Origin,
				Destination: routes[i][j].Destination,
				Departure:   routes[i][j].Departure,
				Arrival:     routes[i][j].Arrival,
				Price:       routes[i][j].Price,
			})
		}

		flights = append(flights, jsonFullFlight{
			Origin:      from,
			Destination: to,
			Patches:     patches,
			Departure:   routes[i][0].Departure,
			Arrival:     routes[i][len(routes[i])-1].Arrival,
		})
	}

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Departure.Before(flights[j].Departure)
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

func initServer(airports Airports, graph *Graph) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/airports", func(w http.ResponseWriter, r *http.Request) {
		airportList(airports, w, r)
	})

	mux.HandleFunc("/flights", func(w http.ResponseWriter, r *http.Request) {
		searchFlights(graph, w, r)
	})

	return http.ListenAndServe(":8080", mux)
}
