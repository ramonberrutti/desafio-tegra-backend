package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type jsonPatch struct {
	Flight      string  `json:"voo"`
	Origin      string  `json:"origem"`
	Destination string  `json:"destino"`
	Departure   string  `json:"saida"`
	Arrival     string  `json:"chegada"`
	Company     string  `json:"operadora"`
	Preco       float32 `json:"preco"`
}

type jsonFullFlight struct {
	Origin      string      `json:"origem"`
	Destination string      `json:"destino"`
	Departure   string      `json:"saida"`
	Arrival     string      `json:"chegada"`
	Patchs      []jsonPatch `json:"trechos"`
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routes)
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
