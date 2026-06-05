package main

import ("fmt"
	"TSP_general/exhaustive"
	"TSP_general/donnees"
	"TSP_general/glouton"
)


func main(){

	var Liste_chemin [][]string
	var meilleur_cout float64
	var meilleur_chemin  []string
	var chemin  []string
	var visitees  []string

	liste_pays := make(map[string]donnees.Coordonnees, 0 )
	liste_pays["FR"] = donnees.Europe()["FR"]
	liste_pays["ES"] = donnees.Europe()["ES"]
	liste_pays["UK"] = donnees.Europe()["UK"]
	liste_pays["IT"] = donnees.Europe()["IT"]
	liste_pays["LI"] = donnees.Europe()["LI"]

	exhaustive.F_exhaustive("FR", "FR", len(liste_pays) , liste_pays, chemin, visitees, &Liste_chemin)

	meilleur_cout = glouton.Cout(Liste_chemin[0], liste_pays)
	meilleur_chemin = Liste_chemin[0]
	var c float64 
	for i:=0; i< len(Liste_chemin); i++{
		c = glouton.Cout(Liste_chemin[i], liste_pays)
		if c < meilleur_cout{
			meilleur_chemin = Liste_chemin[i] 
			meilleur_cout = c
		}
	}
	fmt.Println( meilleur_chemin)
}