package main

import (
	"fmt"
	"TSP_general/opt_glouton"
	"TSP_general/glouton"
	"TSP_general/donnees"
	"time"

)

func main() {

	map_pays := make(map[string]donnees.Coordonnees, 0)
	map_pays["FR"] = donnees.Europe()["FR"]
	map_pays["ES"] = donnees.Europe()["ES"]
	map_pays["UK"] = donnees.Europe()["UK"]
	map_pays["IT"] = donnees.Europe()["IT"]
	map_pays["LI"] = donnees.Europe()["LI"]

	trajet := glouton.Fonc_glouton("FR", map_pays)

	trajet_new := opt_glouton.RetireCroisements
	fmt.Println(trajet)
	fmt.Println(trajet_new)

}
