package main

import (
	"TSP_general/donnees"
	"TSP_general/exhaustive"
	"TSP_general/glouton"
	"fmt"
	"time"
)


func main(){

	start := time.Now()


		map_villes := make(map[string]donnees.Coordonnees, 0)
	map_villes["FR"] = donnees.Europe()["FR"]
	map_villes["ES"] = donnees.Europe()["ES"]
	map_villes["UK"] = donnees.Europe()["UK"]
	map_villes["IT"] = donnees.Europe()["IT"]
	map_villes["LI"] = donnees.Europe()["LI"]
	map_villes["MT"] = donnees.Europe()["MT"]
	map_villes["NL"] = donnees.Europe()["NL"]
	map_villes["NO"] = donnees.Europe()["NO"]
	map_villes["PL"] = donnees.Europe()["PL"]
	map_villes["PT"] = donnees.Europe()["PT"]
	map_villes["SE"] = donnees.Europe()["SE"]

	meilleur_chemin := exhaustive.Exhaustive("FR", map_villes)

	fmt.Println( meilleur_chemin)
	
	duree := time.Since(start)

	cout := glouton.Cout(meilleur_chemin)

	fmt.Println("le programme a pris :", duree, "de cout", cout)

}
