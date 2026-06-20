package branchandbound

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"slices"
)

var min float64

func F_branchandbound_privee(depart string, Liste_chemin [][]string, chemin []string, visites map[string]bool, dico_Villes map[string]donnees.Coordonnees, cout_courant float64) ([][] string){
	if len(chemin) == len(dico_Villes) {
			new_chemin := make([]string, len(chemin)+1)
			copy(new_chemin, chemin)	
			new_chemin = append(new_chemin, depart)
		c := cout_courant + glouton.Distance(donnees.Monde[chemin[len(chemin)-1]],donnees.Monde[depart] )
		if c <= min {
			min = c
			Liste_chemin = append(Liste_chemin, new_chemin)

		}
		chemin = slices.Delete(chemin, len(chemin)-1, len(chemin))
			return Liste_chemin
	} else {
		for ville := range dico_Villes {
			if visites[ville] != true {

				c := cout_courant + glouton.Distance(donnees.Monde[chemin[len(chemin)-1]],donnees.Monde[ville] )
				if min >= c {
					visites[ville] = true
					chemin = append(chemin, ville)
					Liste_chemin = F_branchandbound_privee(depart, Liste_chemin, chemin, visites, dico_Villes, c)
					chemin = chemin[:len(chemin)-1]
					visites[ville] = false
				}

			}
		}
		return Liste_chemin
	}
}

func F_branchandbound(depart string, dico_Villes map[string]donnees.Coordonnees) []string {

	Liste_chemin := make([][]string, 0)
	chemin := []string{depart}
	visites := make(map[string]bool)
	visites[depart] = true
	min = glouton.Cout(glouton.Fonc_glouton(depart, dico_Villes))
	Liste_chemin = F_branchandbound_privee(depart, Liste_chemin, chemin, visites, dico_Villes,0)


	var c float64 

	if len(Liste_chemin) == 0 {
		return glouton.Fonc_glouton(depart, dico_Villes)
	}else {
		meilleur_chemin := Liste_chemin[0]
			for  i:=0; i< len(Liste_chemin); i++{
		c = glouton.Cout(Liste_chemin[i])
		if c < min{
			meilleur_chemin = Liste_chemin[i] 
			min = c
		}

	}

return meilleur_chemin
}}
	

