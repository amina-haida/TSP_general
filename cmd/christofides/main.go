package main

import (
	"TSP_general/christofides"
	"TSP_general/donnees"
	"TSP_general/dynamique"
	"TSP_general/glouton"
	"TSP_general/kruskal"

	"fmt"
)
	
func main(){


	europe := donnees.Europe()

	chemin_min := christofides.Christofides("FR", europe)
	fmt.Println("Le chemin  est : ", chemin_min)
	duree := time.Since(start)

	fmt.Println("le programme a pris :", duree)
}