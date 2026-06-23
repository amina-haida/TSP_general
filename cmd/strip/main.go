package main

import (
	"fmt"
	"TSP_general/donnees"
	"TSP_general/strip"
)

func main(){
	trajet := strip.Strip(donnees.Europe())
	fmt.Println(trajet)
}