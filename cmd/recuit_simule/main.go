package main

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"TSP_general/recuit_simule"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().UnixNano())

	trajet := glouton.Fonc_glouton("FR", donnees.Monde)

	recuit_simule.Recuit(trajet, 5)

}