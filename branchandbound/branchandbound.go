package branchandbound

import(
	"TSP_general/donnees"
	"slices"
	"TSP_general/glouton")



var Liste_chemin [][]string
var chemin []string
var visites map[string]bool
var map_pays map[string]donnees.Coordonnees
var min float64


func F_branchandbound_privee( depart string,  Liste_chemin *[][]string, chemin []string,  visites map[string]bool ) {
		if len(chemin) == len(map_pays) {
			chemin = append(chemin, depart)
			new_chemin := make([]string, len(chemin))
			copy(new_chemin, chemin)   
			*Liste_chemin = append(*Liste_chemin, new_chemin)
			c:= glouton.Cout(chemin)
			if c < min {
				min = c
			}
			return
		}else if len(map_pays) < 2  {
				return  
		}else {
			for pays := range map_pays {
				if !visites[pays]{
					chemin = append(chemin, pays)
					c := glouton.Cout(chemin)
					if min > c{
						visites[pays]=true
						F_exhaustive_privee(depart,  Liste_chemin , chemin , visites )
						chemin = slices.Delete(chemin, len(chemin)-1, len(chemin))
						visites[pays]=false
					}
					
				}
			}
		}
	}


func F_branchandbound( depart string, lp map[string]donnees.Coordonnees) [][]string{

	map_pays = lp


	Liste_chemin = make([][]string, 0)
	chemin := []string{depart}
	visites := make(map[string]bool)
	visites[depart]=true
	min = glouton.Cout(glouton.Fonc_glouton(depart, lp))
	F_exhaustive_privee(depart,  &Liste_chemin , chemin , visites)
	return Liste_chemin
	
}
