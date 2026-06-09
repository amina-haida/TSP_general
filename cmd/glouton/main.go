package main

import (
	"fmt"
	"TSP_general/glouton"
	"TSP_general/donnees"
	"time"
)

func main() {

	start := time.Now()

	map_pays := make(map[string]donnees.Coordonnees, 0)
	map_pays["FR"] = donnees.Europe()["FR"]
	map_pays["ES"] = donnees.Europe()["ES"]
	map_pays["UK"] = donnees.Europe()["UK"]
	map_pays["IT"] = donnees.Europe()["IT"]
	map_pays["LI"] = donnees.Europe()["LI"]

	tour := glouton.Fonc_glouton("FR", map_pays)

	fmt.Println(tour)
	fmt.Println(glouton.Cout(tour))

	duree := time.Since(start)

	fmt.Println("Le programme a duré :", duree)
	
}
