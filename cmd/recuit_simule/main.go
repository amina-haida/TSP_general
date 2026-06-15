package main

import(
	"TSP_general/donnees"
	"TSP_general/glouton"
	"TSP_general/opt_glouton"
	"TSP_general/recuit_simule"
)


func main() {
	trajet := glouton.Fonc_glouton("FR", donnees.Pays)

	trajetOpt := opt_glouton.RetireCroisements(trajet, donnees.Pays)

	recuit_simule.Recuit(trajetOpt, 20)

}