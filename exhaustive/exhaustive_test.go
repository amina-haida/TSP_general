package exhaustive

import (
	"testing"
	"TSP_general/donnees"
)

func Test_F_exhaustive(t *testing.T){
	var Liste_chemin [][]string



	dico_villes := make(map[string]donnees.Coordonnees, 0 )
	dico_villes["FR"] = donnees.Europe()["FR"]
	dico_villes["ES"] = donnees.Europe()["ES"]
	dico_villes["UK"] = donnees.Europe()["UK"]
	dico_villes["IT"] = donnees.Europe()["IT"]
	dico_villes["LI"] = donnees.Europe()["LI"]


	Liste_chemin = F_exhaustive("FR", dico_villes)
	if len(Liste_chemin) != 24 {
		 t.Fatalf(
            "nombre de chemins incorrect : obtenu %d, attendu %d",
            len(Liste_chemin),
            24,
        )
	}
	
	for index := range Liste_chemin{
		if Liste_chemin[index][0] != "FR" || Liste_chemin[index][len(Liste_chemin[index])-1]!="FR" || len(Liste_chemin[index])!=6 {
			t.Errorf("Ce n'est pas une boucle")
		}
	}
}
