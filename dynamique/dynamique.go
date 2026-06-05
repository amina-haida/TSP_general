package dynamique

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"math"
)



func Makelist( donnees map[string]donnees.Coordonnees) ([]string){
	var new []string
	for pays,_ := range donnees {
		new = append(new, pays)
	}
	return new
}

func Makeint(ensemble  map[string]string, liste []string)(string){
	var nombre string
	for index := range liste{
		if ensemble[liste[index]] == "1" {
			nombre = nombre + "1"
		}else {
			nombre = nombre + "0"
		}

	}
		return nombre 
}



// faudra initiliaser visitees avec que des "0"

func Cout_dynamique( depart string, arrivee string, visitees map[string]string, memo_cout map[string]float64, liste_pays []string){
	if len(visitees) == 1{
		visitees[arrivee] = "1"
		nombre := Makeint(visitees, liste_pays)
		memo_cout[nombre] = glouton.Distance(depart, arrivee, donnees.Pays)
	}else {
		min_cout := math.Inf(1)

		for _, pays := range visitees {
			visitees[pays]="0"
			nombre1 := Makeint(visitees, liste_pays)
			calcul := memo_cout[nombre1]
			glouton.Distance(pays, arrivee, donnees.Pays)
			
			if pays != depart && calcul < min_cout {

				min_cout= calcul

			}
			visitees[pays]="1"
		}
		nombre2 := Makeint(visitees, liste_pays)
		memo_cout[nombre2]= min_cout
	}
}

