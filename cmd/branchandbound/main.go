package main

import (
	"TSP_general/branchandbound"
	"TSP_general/donnees"
	"TSP_general/glouton"
	"fmt"
	"time"
)

func main (){


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

		
	start := time.Now()


	chemin_min := branchandbound.F_branchandbound("FR",map_villes)

	duree := time.Since(start)
	cout := glouton.Cout(chemin_min, donnees.Europe())

	fmt.Println("Le chemin optimal est ", chemin_min, "calculé en", duree, "de cout", cout)
}