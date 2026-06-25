package main 

import ("TSP_general/christofides"
"TSP_general/donnees"
"time"
"fmt")
	
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
	

		chemin_min := christofides.Christofides("FR",map_villes)
		fmt.Println("Le chemin  est : ", chemin_min)
		duree := time.Since(start)

		fmt.Println("le programme a pris :", duree)
}