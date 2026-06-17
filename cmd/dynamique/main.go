package main

import (
	"fmt"
	"TSP_general/dynamique"
	"TSP_general/donnees"
	"time"

)

func main(){

	start := time.Now()


	map_pays := make(map[string]donnees.Coordonnees, 0)
	map_pays["FR"] = donnees.Europe()["FR"]
	map_pays["ES"] = donnees.Europe()["ES"]
	map_pays["UK"] = donnees.Europe()["UK"]
	map_pays["IT"] = donnees.Europe()["IT"]
	map_pays["LI"] = donnees.Europe()["LI"]
	map_pays["MT"] = donnees.Europe()["MT"]
	map_pays["NL"] = donnees.Europe()["NL"]
	map_pays["NO"] = donnees.Europe()["NO"]
	map_pays["PL"] = donnees.Europe()["PL"]
	map_pays["PT"] = donnees.Europe()["PT"]
	map_pays["SE"] = donnees.Europe()["SE"]

	dynamique.Memo_pred = make(map[dynamique.Etat]string)
	dynamique.Memo_cout = make(map[dynamique.Etat]float64)

	dynamique.Liste_pays = dynamique.Makelist(map_pays)

	etat := dynamique.Etat{ Visites : "11111111111", Arrive : "FR"}

	cout_min := dynamique.Cout_dynamique(etat,"FR")

	var chemin_min []string

	for i :=0; i < len(dynamique.Makelist(map_pays)); i++ {

		pays := dynamique.Memo_pred[etat]
		chemin_min = append(chemin_min, pays)
		etat = dynamique.Enlever(etat, pays)
}
		chemin_min = append(chemin_min, "FR")
		fmt.Println("Le chemin optimal est : ", chemin_min, " avec un coût de : ", cout_min)
		duree := time.Since(start)

		fmt.Println("le programme a pris :", duree)
	}

