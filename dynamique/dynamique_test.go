package dynamique

import (
	"testing"
	"TSP_general/donnees"
	"slices"
	"TSP_general/glouton"
)




func Test_Makeliste(t *testing.T){
	dico_villes := map[string]donnees.Coordonnees {
		"FR": {Latitude : 46.227638, Longitude: 2.213749},
		"GA": {Latitude :-0.803689, Longitude:  11.609444},
		"GB": {Latitude : 55.378051, Longitude: -3.435973},
	}
	res := []string{"FR", "GA","GB"}
	if !slices.Equal(res, donnees.Makelist(dico_villes)){
		    t.Fatalf("attendu pas ce qu'on a eu")
    }

}

func Test_longueur(t *testing.T) {
    etat := Etat{
        Pas_Visitees: "101",
        Arrivee:  "FR",
    }

    res := longueur(etat)
    if res != 2 {
        t.Fatalf("attendu 2, obtenu %d", res)
    }
}


func Test_paysvisites(t *testing.T) {
    etat := Etat{
        Pas_Visitees: "101",
        Arrivee:  "FR",
    }
	dico_villes := map[string]donnees.Coordonnees {
		"FR": {Latitude : 46.227638, Longitude: 2.213749},
		"GA": {Latitude :-0.803689, Longitude:  11.609444},
		"GB": {Latitude : 55.378051, Longitude: -3.435973},
	}

	res := []string{"FR", "GB"}
	Liste_villes = donnees.Makelist(dico_villes)
	if !slices.Equal(res, villeavisitees(etat)){
		    t.Fatalf("attendu pas ce qu'on a eu")
    }
	}


func Test_enlever(t *testing.T){
	etat := Etat{
        Pas_Visitees: "101",
        Arrivee:  "FR",
    }
	dico_villes := map[string]donnees.Coordonnees {
		"FR": {Latitude : 46.227638, Longitude: 2.213749},
		"GA": {Latitude :-0.803689, Longitude:  11.609444},
		"GB": {Latitude : 55.378051, Longitude: -3.435973},
	}
	Liste_villes = donnees.Makelist(dico_villes)
	etat2 := Etat{ Pas_Visitees : "001", Arrivee : "FR"}

	if Enlever(etat, "FR") != etat2{
		t.Fatalf("attendu pas ce qu'on a eu")
	}

	
}
func Test_enlever_arrive(t *testing.T) {
    Liste_villes = []string{"FR","GA","GB"}

    etat := Etat{
        Pas_Visitees: "101",
        Arrivee: "FR",
    }

    obtenu := Enlever(etat, "GB")

    attendu := Etat{
        Pas_Visitees: "100",
        Arrivee: "GB",
    }

    if obtenu != attendu {
        t.Fatalf("attendu %+v, obtenu %+v", attendu, obtenu)
    }
}
func Test_Cout_dynamique(t *testing.T){

	Memo_cout = make(map[Etat]float64)
	Memo_pred = make(map[Etat]string)
	etat := Etat{
        Pas_Visitees: "000",
        Arrivee:  "FR",
    }

	dico_villes2 := map[string]donnees.Coordonnees {
		"FR": {Latitude : 46.227638, Longitude: 2.213749},
		"GA": {Latitude :-0.803689, Longitude:  11.609444},
		"GB": {Latitude : 55.378051, Longitude: -3.435973},
	}
	Liste_villes = donnees.Makelist(dico_villes2)

	res:= Cout_dynamique(etat, "GA")

	res_attendu := glouton.Distance(donnees.Monde["GA"], donnees.Monde["FR"])

	if res != res_attendu {
		t.Fatalf("attendu pas ce qu'on a eu")
	}


}

func Test_Pred(t *testing.T) {

    Memo_cout = make(map[Etat]float64)
    Memo_pred = make(map[Etat]string)

    Liste_villes = []string{"FR","GA","GB"}

    etat := Etat{
        Pas_Visitees: "011",
        Arrivee: "FR",
    }

    Cout_dynamique(etat, "FR")

    pred := Memo_pred[etat]

    if pred != "GA" {
		t.Fatalf("attendu GA, obtenu %s", pred)
	}
}