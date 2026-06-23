package main

import (
	"fmt"
	"TSP_general/donnees"
	"TSP_general/glouton"
	"TSP_general/recuit_simule"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	trajet := glouton.Fonc_glouton("FR", donnees.Monde)
	trajet_new := recuit_simule.Recuit("FR", donnees.Monde)
	
	fmt.Println(trajet)
	fmt.Println(trajet_new)

}