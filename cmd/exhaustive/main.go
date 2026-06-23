package main

import (
	"TSP_general/donnees"
	"TSP_general/exhaustive"
	"fmt"
	"time"
)


func main(){

	start := time.Now()

depart := "FR"
	dicoVilles := donnees.Ville_aléatoire(11)
	_ = exhaustive.Exhaustive(depart, dicoVilles)

	duree := time.Since(start)

	fmt.Println("le programme a pris :", duree)

}
