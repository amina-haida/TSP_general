package main

import (
	"fmt"
	"TSP_general/glouton"
	"TSP_general/donnees"
)

func main() {

	europe := donnees.Europe()

	tour := glouton.Fonc_glouton("FR", europe)

	fmt.Println(tour)
	fmt.Println(glouton.Cout(tour))
}
