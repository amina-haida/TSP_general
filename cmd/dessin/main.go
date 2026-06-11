package main

import (
	"TSP_general/dessin"
	"TSP_general/donnees"
	"TSP_general/glouton"
	"TSP_general/opt_glouton"
)

func main() {
	europe := donnees.Europe()
	trajet := glouton.Fonc_glouton("FR", europe)
	
	trajet_new := opt_glouton.RetireCroisements(trajet, europe)


	dessin.Dessiner_trajet(europe, trajet_new)

}