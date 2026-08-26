package main

import (
	"TSP_general/branchandbound"
	"TSP_general/donnees"
	"TSP_general/glouton"
	"fmt"
	"time"
)

func main (){
	map_pays := donnees.Ville_aleatoire(14)	
	start := time.Now()



		chemin_min := branchandbound.F_branchandbound("PT",map_pays)
		fmt.Println("Le chemin optimal est : ", chemin_min)
		duree := time.Since(start)

	cout := glouton.Cout(chemin_min, donnees.Europe())

	fmt.Println("Le chemin optimal est ", chemin_min, "calculé en", duree, "de cout", cout)
}