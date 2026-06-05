package exhaustive

import(
	"TSP_general/donnees"
	"slices")


func F_exhaustive( depart string, arrivee string, longueur int, liste_pays map[string]donnees.Coordonnees, chemin []string , visitees []string, Liste_chemin *[][]string) {

	if len(chemin) == longueur {
		chemin = append(chemin, arrivee)
		new_chemin := make([]string, len(chemin))
		copy(new_chemin, chemin)   
		*Liste_chemin = append(*Liste_chemin, new_chemin)
		return
	}else if longueur < 2 ||longueur > len(liste_pays) {
			return  
		}else if len(chemin) == 0 {
		visitees = append(visitees, depart)
		chemin = append(chemin, depart)
		F_exhaustive(depart, arrivee, longueur, liste_pays,chemin,visitees, Liste_chemin )
	}else {
		for pays := range liste_pays {
			if !slices.Contains(visitees,pays){
				chemin = append(chemin, pays)
				visitees = append(visitees, pays)
				F_exhaustive(depart, arrivee, longueur, liste_pays,chemin,visitees, Liste_chemin )
				chemin = slices.Delete(chemin, len(chemin)-1, len(chemin))
				visitees = slices.Delete(visitees, len(visitees)-1, len(visitees))
			}
		}
	}
}

