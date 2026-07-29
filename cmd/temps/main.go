package main

import (
	"TSP_general/donnees"
	"TSP_general/temps"

	"math/rand"
	"time"
)

func main(){
 rand.Seed(time.Now().UnixNano())
    temps.Comparaison_heuristiques(len(donnees.Monde))

}
