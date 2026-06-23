package main

import (
	"fmt"
	"TSP_general/opt_glouton"
	"TSP_general/glouton"
	"TSP_general/donnees"
)

func main() {

	trajet := glouton.Fonc_glouton("FR", donnees.Monde)

	trajet_new := opt_glouton.Fonc_optGlouton("FR", donnees.Monde)
	
	fmt.Println(trajet)
	fmt.Println(trajet_new)

}
