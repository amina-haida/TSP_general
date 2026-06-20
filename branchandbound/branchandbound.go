package branchandbound

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"math"
)

var min float64
var meilleur_chemin []string


func F_branchandbound_privee(depart string, chemin []string, visites map[string]bool, dico_Villes map[string]donnees.Coordonnees, cout_courant float64) {
	if len(chemin) == len(dico_Villes) {
			new_chemin := make([]string, len(chemin)+1)
			copy(new_chemin, chemin)	
			new_chemin[len(chemin)]=depart
		c := cout_courant + glouton.Distance(dico_Villes[chemin[len(chemin)-1]],dico_Villes[depart] )
		if c <= min {
			min = c
			meilleur_chemin = new_chemin

		}
			return 
	} else {
		for ville := range dico_Villes {
			if visites[ville] != true {

				c := cout_courant + glouton.Distance(dico_Villes[chemin[len(chemin)-1]], dico_Villes[ville] )
				if min >= c {
					visites[ville] = true
					chemin = append(chemin, ville)
					 F_branchandbound_privee(depart, chemin, visites, dico_Villes, c)
					chemin = chemin[:len(chemin)-1]
					visites[ville] = false
				}

			}
		}
		return 
	}
}

func F_branchandbound(depart string, dico_Villes map[string]donnees.Coordonnees) []string {
	chemin := []string{depart}
	meilleur_chemin = glouton.Fonc_glouton(depart, dico_Villes)
	visites := make(map[string]bool)
	visites[depart] = true
	min = math.MaxFloat64
	F_branchandbound_privee(depart, chemin, visites, dico_Villes,0)

return meilleur_chemin
}
	

