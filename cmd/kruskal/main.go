package main

import (
	"TSP_general/kruskal"
	"TSP_general/donnees"
	"time"
	"fmt"
)

func main (){
		start := time.Now()


	map_pays := make(map[string]donnees.Coordonnees, 0)
	map_pays["FR"] = donnees.Europe()["FR"]
	map_pays["ES"] = donnees.Europe()["ES"]
	map_pays["UK"] = donnees.Europe()["UK"]
	map_pays["IT"] = donnees.Europe()["IT"]
	map_pays["LI"] = donnees.Europe()["LI"]

	chemin_min := kruskal.Kruskal(map_pays)
		fmt.Println("Le chemin optimal est : ", chemin_min)

		duree := time.Since(start)
		fmt.Println("le programme a pris :", duree)

}