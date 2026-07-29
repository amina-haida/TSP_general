package main

import (
	"TSP_general/christofides"
	"TSP_general/donnees"
	"TSP_general/dynamique"
	"TSP_general/glouton"
	"TSP_general/kruskal"

	"fmt"
)
	
func main(){


		map_villes := donnees.Ville_aleatoire(11)
depart :=""

for ville := range map_villes{
	depart = ville
}
	

		chemin_min_c := christofides.Christofides(depart,map_villes)
		chemin_min_k := kruskal.Kruskal_TSP(depart,map_villes)
		chemin := dynamique.TSP_dynamique(depart,map_villes)

		cout_c := glouton.Cout(chemin_min_c, map_villes)
			cout_k := glouton.Cout(chemin_min_k,map_villes)
				cout:= glouton.Cout(chemin,map_villes)
			
		fmt.Println("Le chemin  est : ", chemin,"de cout", cout, "kruskal : ", chemin_min_k,"de cout",cout_k, "christofides", chemin_min_c, "de cout",cout_c)

}