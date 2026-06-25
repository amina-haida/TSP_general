package main

import (
	"fmt"
	"TSP_general/donnees"
	"TSP_general/strip"
)

func main(){
	trajet := strip.MeilleurStrip(donnees.Europe(), "FR")
	fmt.Println(trajet)
}