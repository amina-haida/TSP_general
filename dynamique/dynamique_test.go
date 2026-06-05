package dynamique

import ("testing"
		"TSP_general/donnees"
	"slices")




func Test_Makeliste(t *testing.T){
	donnees := map[string]donnees.Coordonnees {
		"FR": {Latitude : 46.227638, Longitude: 2.213749},
		"GA": {Latitude :-0.803689, Longitude:  11.609444},
		"GB": {Latitude : 55.378051, Longitude: -3.435973},
	}
	res := []string{"FR", "GA","GB"}
	if !slices.Equal(res, Makelist(donnees)){
		    t.Fatalf("attendu pas ce qu'on a eu")
    }

}

func Test_longueur(t *testing.T) {
    etat := Etat{
        visites: "101",
        arrive:  "FR",
    }

    res := longueur(etat)
    if res != 2 {
        t.Fatalf("attendu 2, obtenu %d", res)
    }
}



func Test_pays(t *testing.T) {
    etat := Etat{
        visites: "101",
        arrive:  "FR",
    }
	donnees := map[string]donnees.Coordonnees {
		"FR": {Latitude : 46.227638, Longitude: 2.213749},
		"GA": {Latitude :-0.803689, Longitude:  11.609444},
		"GB": {Latitude : 55.378051, Longitude: -3.435973},
	}
	var liste_pays []string
	res := []string{"FR", "GB"}
	liste_pays =Makelist(donnees)
	if !slices.Equal(res, pays(etat, liste_pays)){
		    t.Fatalf("attendu pas ce qu'on a eu")
    }
	}


func Test_enlever(t *testing.T){
	
}

