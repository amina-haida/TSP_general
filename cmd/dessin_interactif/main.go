package main

import(
	"TSP_general/donnees"
	"TSP_general/dessin_interactif"
)


func main() {
	europe := donnees.Europe()

	dessininteractif.Dessin_points(europe)

}
