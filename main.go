package main

import (
	"log"
)

func main() {
	var graph Graph

	// Move loaders to another place!
	airports, err := LoadAirportsFromFile("data/aeroportos.json")
	if err != nil {
		log.Fatal(err)
	}

	for i := range airports {
		node := Node(airports[i])
		graph.AddNode(airports[i].Name, &node)
	}

	flightsUber, err := LoadFlightsFromFile("data/uberair.csv", "UberAir")
	if err != nil {
		log.Fatal(err)
	}

	flights99, err := LoadFlightsFromFile("data/99planes.json", "99Planes")
	if err != nil {
		log.Fatal(err)
	}

	flights := append(flightsUber, flights99...)

	for i := range flights {
		edge := Edge(flights[i])
		graph.AddEdge(flights[i].Origin, flights[i].Destination, &edge)
	}

	// routes := graph.FoundRoute("VCP", "BEL", func(nameFrom string, edge Edge) bool {
	// 	if edge.Departure.Year() == 2019 && edge.Departure.Month() == 2 && edge.Departure.Day() == 17 {
	// 		return true
	// 	}

	// 	return false
	// }, func(nameFrom string, edgeFrom, edge Edge) bool {
	// 	diff := edge.Departure.Sub(edgeFrom.Arrival)
	// 	if diff >= 0 && diff <= time.Hour*12 {
	// 		return true
	// 	}

	// 	return false
	// })

	// // log.Println(routes)

	// b, err := json.Marshal(routes)

	// log.Println(string(b))

	log.Fatal(initServer(airports, &graph))

}
