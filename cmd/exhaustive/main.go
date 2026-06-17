package main

import ("fmt"
	"TSP_general/exhaustive"
	"TSP_general/donnees"
	"TSP_general/glouton"
	"time"
)


func main(){

	start := time.Now()
	var Liste_chemin [][]string
	var meilleur_cout float64
	var meilleur_chemin  []string
	
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

	Liste_chemin=exhaustive.F_exhaustive("FR", map_pays)

	meilleur_cout = glouton.Cout(Liste_chemin[0])
	meilleur_chemin = Liste_chemin[0]
	var c float64 
	for i:=0; i< len(Liste_chemin); i++{
		c = glouton.Cout(Liste_chemin[i])
		if c < meilleur_cout{
			meilleur_chemin = Liste_chemin[i] 
			meilleur_cout = c
		}
	}
	fmt.Println( meilleur_chemin)
	
	duree := time.Since(start)

	fmt.Println("le programme a pris :", duree)

}
