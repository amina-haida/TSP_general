package main

import (
	"fmt"
	"TSP_general/troncon"
)

var chemin []string

func main() {

	troncon.Reseau_ferre = []troncon.Troncon{
		{Depart : "1", Arrivee:  "2"},
		{Depart :"2", Arrivee: "3"},
		{Depart :"2", Arrivee: "4"},
		{Depart :"1",Arrivee:  "5"},
		{Depart : "1",Arrivee:  "6"},
	}
	
	troncon.Troncons_rec("1", "1")
	
	troncon.Troncons_rec2("1", "1", &chemin)

	for index := range chemin {
		fmt.Println(chemin[index])
	}
	
	troncon.Chercher_ville("1", "1", "4", &chemin)

	for _, ville := range chemin {
		fmt.Println(ville)
	}
}
