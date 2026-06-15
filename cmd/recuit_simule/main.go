package main

import(
	"fmt"
	"TSP_general/glouton"
	"TSP_general/donnees"
	"TSP_general/recuit_simule"
)


func main() {
	trajet := glouton.Fonc_glouton("FR", donnees.Pays)

	trajetMin := recuit_simule.Recuit(trajet, 10)

	fmt.Println(trajetMin)
}