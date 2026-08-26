package main

import (
	"TSP_general/donnees"
	"TSP_general/exhaustive"
	"TSP_general/glouton"
	"fmt"
	"time"
)


func main(){



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
	
	start := time.Now()



		chemin_min := exhaustive.Exhaustive("PT",map_pays)
		fmt.Println("Le chemin optimal est : ", chemin_min)
		duree := time.Since(start)
	cout := glouton.Cout(chemin_min, map_pays)
		fmt.Println("le programme a pris :", duree, "de cout",cout)

}
