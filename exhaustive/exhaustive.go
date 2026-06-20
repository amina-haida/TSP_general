package exhaustive

import(
	"TSP_general/donnees"
	"slices"
"TSP_general/glouton")


	
var Liste_chemin [][]string
var chemin []string
var visites map[string]bool


func Recherche_chemins( depart string,  Liste_chemin *[][]string, chemin []string,  visites map[string]bool, dico_villes map[string]donnees.Coordonnees ) {
		if len(chemin) == len(dico_villes) {
			chemin = append(chemin, depart)
			new_chemin := make([]string, len(chemin))
			copy(new_chemin, chemin)   
			*Liste_chemin = append(*Liste_chemin, new_chemin)
			return  
		}else {
			for pays := range dico_villes {
				if !visites[pays]{
					chemin = append(chemin, pays)
					visites[pays]=true
					Recherche_chemins(depart,  Liste_chemin , chemin , visites, dico_villes )
					chemin = slices.Delete(chemin, len(chemin)-1, len(chemin))
					visites[pays]=false
				}
			}
		}
	}


func F_exhaustive( depart string, dico_villes map[string]donnees.Coordonnees) [][]string{


	Liste_chemin = make([][]string, 0)
	chemin := []string{depart}
	visites := make(map[string]bool)
	visites[depart]=true

	Recherche_chemins(depart,  &Liste_chemin , chemin , visites, dico_villes)
	return Liste_chemin
	
}


func Exhaustive(depart string, dico_villes map[string]donnees.Coordonnees)  []string{

	var Liste_chemin [][]string
	var meilleur_cout float64
	var meilleur_chemin  []string
	
	Liste_chemin=F_exhaustive("FR", dico_villes)

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
	return meilleur_chemin
}