package main

import (
	"TSP_general/dynamique"
	"TSP_general/donnees"
	"math"

)

func main(){

	map_pays := make(map[string]donnees.Coordonnees, 0 )
	map_pays["FR"] = donnees.Europe()["FR"]
	map_pays["ES"] = donnees.Europe()["ES"]
	map_pays["UK"] = donnees.Europe()["UK"]
	map_pays["IT"] = donnees.Europe()["IT"]
	map_pays["LI"] = donnees.Europe()["LI"]

	memo_cout := make(map[dynamique.Etat] float64)
	memo_pred := make(map[dynamique.Etat] string)

	etat := dynamique.Etat{ Visites : "00000", Arrive : "FR"}

	dynamique.Cout_dynamique(etat,"FR", dynamique.Makelist(map_pays))

	cout_min := math.Inf(1)
	chemin_min := memo_pred[etat]

	for etat, cout := range memo_cout {
		if cout < cout_min {
			cout_min = cout
			chemin_min = memo_pred[etat]
		}		
	println("Le chemin optimal est : ", chemin_min, " avec un coût de : ", cout_min)
		}
	}