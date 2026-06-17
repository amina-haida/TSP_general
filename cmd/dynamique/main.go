package main

import (
	"fmt"
	"TSP_general/dynamique"
	"TSP_general/donnees"
	"time"

)

func main(){

	start := time.Now()


	map_pays := make(map[string]donnees.Coordonnees, 0)
	map_pays["FR"] = donnees.Europe()["FR"]
	map_pays["ES"] = donnees.Europe()["ES"]
	map_pays["UK"] = donnees.Europe()["UK"]
	map_pays["IT"] = donnees.Europe()["IT"]
	map_pays["LI"] = donnees.Europe()["LI"]
	map_pays["MT"] = donnees.Europe()["MT"]
	map_pays["NL"] = donnees.Europe()["NL"]
	map_pays["NO"] = donnees.Europe()["NO"]
	map_pays["PL"] = donnees.Europe()["PL"]
	map_pays["PT"] = donnees.Europe()["PT"]
	map_pays["SE"] = donnees.Europe()["SE"]

		chemin_min := dynamique.TSP_dynamique(map_pays, "FR")
		fmt.Println("Le chemin optimal est : ", chemin_min)
		duree := time.Since(start)

		fmt.Println("le programme a pris :", duree)
	}

