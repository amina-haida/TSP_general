package main

import (
	"fmt"
	"TSP_general/opt_glouton"
	"TSP_general/glouton"
	"TSP_general/donnees"
)

func main() {

	trajet := glouton.Fonc_glouton("FR", donnees.Pays)

	trajet_new := opt_glouton.RetireCroisements(trajet, donnees.Pays)
	fmt.Println(trajet)
	fmt.Println(trajet_new)

}
