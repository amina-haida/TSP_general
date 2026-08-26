package main

import (
	"TSP_general/temps"
	"TSP_general/donnees"
	"math/rand"
	"time"
)

func main(){
 rand.Seed(time.Now().UnixNano())
    temps.Comparaison_heuristiques(len(donnees.Monde))

}
