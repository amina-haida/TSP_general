package main

import (
	"TSP_general/dessin"
	"TSP_general/donnees"
	"TSP_general/glouton"
)

func main() {
	europe := donnees.Europe()
	trajet := glouton.Fonc_glouton("FR", europe)

	dessin.Dessiner_points(europe)

	dessin.Dessiner_trajet(europe, trajet)

}