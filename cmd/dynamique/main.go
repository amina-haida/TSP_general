package main

import (
	"fmt"
	"TSP_general/dynamique"
	"TSP_general/donnees"

)

func main(){

	map_pays := make(map[string]donnees.Coordonnees, 0)
	map_pays["FR"] = donnees.Europe()["FR"]
	map_pays["ES"] = donnees.Europe()["ES"]
	map_pays["UK"] = donnees.Europe()["UK"]
	map_pays["IT"] = donnees.Europe()["IT"]
	map_pays["LI"] = donnees.Europe()["LI"]

	dynamique.Memo_pred = make(map[dynamique.Etat]string)
	dynamique.Memo_cout = make(map[dynamique.Etat]float64)


	etat := dynamique.Etat{ Visites : "11111", Arrive : "FR"}

	cout_min := dynamique.Cout_dynamique(etat,"FR", dynamique.Makelist(map_pays))

	var chemin_min []string

	for i :=0; i < len(dynamique.Makelist(map_pays)); i++ {

		pays := dynamique.Memo_pred[etat]
		chemin_min = append(chemin_min, pays)
		etat = dynamique.Enlever(etat, pays, dynamique.Makelist(map_pays))


		
	
		}
		chemin_min = append(chemin_min, "FR")
		fmt.Println("Le chemin optimal est : ", chemin_min, " avec un coût de : ", cout_min)
	}

