package main

import (
	"fmt"
	"TSP_general/donnees"
	"TSP_general/strip"
)

func main(){
	trajet := strip.MeilleurStrip("FR", donnees.Europe())
	fmt.Println(trajet)
}