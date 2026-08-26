package main

import (
	"TSP_general/kruskal"
	"TSP_general/donnees"
	"TSP_general/glouton"
	"time"
	"fmt"
)

func main (){
	start := time.Now()


	europe := donnees.Europe()
	
	chemin_min := kruskal.Kruskal_TSP("FR", europe)
	fmt.Println("Le chemin optimal est : ", chemin_min)

	duree := time.Since(start)
	cout := glouton.Cout(chemin_min, europe)
	fmt.Println("le programme a pris :", duree,"de coût", cout)

}