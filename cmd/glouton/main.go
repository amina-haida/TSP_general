package main

import (
	"fmt"
	"math"
	"TSP_general/glouton"
	"TSP_general/donnees"
)

func main() {

	min := math.Inf(1)
	max := 0.0
	somme := 0.0
	nb := 0

	for pays1 := range donnees.Pays {
		for pays2 := range donnees.Pays {

			if pays1 >= pays2 {
				continue
			}

			d := glouton.Distance(pays1, pays2, donnees.Pays)

			if d < min {
				min = d
			}

			if d > max {
				max = d
			}

			somme += d
			nb++
		}
	}

	moyenne := somme / float64(nb)

	fmt.Println(min)
	fmt.Println(max)
	fmt.Println(moyenne)


	europe := donnees.Europe()

	tour := glouton.Fonc_glouton("FR", europe)

	fmt.Println(tour)
	fmt.Println(glouton.Cout(tour, europe))
}
