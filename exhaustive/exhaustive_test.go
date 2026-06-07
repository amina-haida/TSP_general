package exhaustive

import (
	"testing"
	"TSP_general/donnees"
)

func Test_F_exhaustive(t *testing.T){
	var Liste_chemin [][]string
	var chemin  []string
	var visitees  []string

	liste_pays := make(map[string]donnees.Coordonnees, 0 )
	liste_pays["FR"] = donnees.Europe()["FR"]
	liste_pays["ES"] = donnees.Europe()["ES"]
	liste_pays["UK"] = donnees.Europe()["UK"]
	liste_pays["IT"] = donnees.Europe()["IT"]
	liste_pays["LI"] = donnees.Europe()["LI"]

	expected := 24

	F_exhaustive("FR", "FR", len(liste_pays) , liste_pays, chemin, visitees, &Liste_chemin)
	if len(Liste_chemin) != expected {
		t.Errorf("il n'y a pas tous les chemins")
	}
	
	for index := range Liste_chemin{
		if Liste_chemin[index][0] != "FR" || Liste_chemin[index][len(Liste_chemin[index])-1]!="FR" || len(Liste_chemin[index])!=6 {
			t.Errorf("Ce n'est pas une boucle")
		}
	}
}
