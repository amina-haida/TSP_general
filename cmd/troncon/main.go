package main

import (
	"fmt"
	"TSP_general/troncon"
)

var chemin []string

func main() {

	troncon.Reseau_ferre = []troncon.Troncon{
		{"1", "2"},
		{"2", "3"},
		{"2", "4"},
		{"1", "5"},
		{"1", "6"},
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
