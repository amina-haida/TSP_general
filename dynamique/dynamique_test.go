package dynamique

import (
	"testing"
	"TSP_general/donnees"
	"slices"
	"TSP_general/glouton")




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
        Visites: "101",
        Arrive:  "FR",
    }

    res := longueur(etat)
    if res != 2 {
        t.Fatalf("attendu 2, obtenu %d", res)
    }
}


func Test_paysvisites(t *testing.T) {
    etat := Etat{
        Visites: "101",
        Arrive:  "FR",
    }
	donnees := map[string]donnees.Coordonnees {
		"FR": {Latitude : 46.227638, Longitude: 2.213749},
		"GA": {Latitude :-0.803689, Longitude:  11.609444},
		"GB": {Latitude : 55.378051, Longitude: -3.435973},
	}

	res := []string{"FR", "GB"}
	Liste_pays =Makelist(donnees)
	if !slices.Equal(res, paysvisites(etat)){
		    t.Fatalf("attendu pas ce qu'on a eu")
    }
	}


func Test_enlever(t *testing.T){
	etat := Etat{
        Visites: "101",
        Arrive:  "FR",
    }
	donnees := map[string]donnees.Coordonnees {
		"FR": {Latitude : 46.227638, Longitude: 2.213749},
		"GA": {Latitude :-0.803689, Longitude:  11.609444},
		"GB": {Latitude : 55.378051, Longitude: -3.435973},
	}
	Liste_pays = Makelist(donnees)
	etat2 := Etat{ Visites : "001", Arrive : "FR"}

	if Enlever(etat, "FR") != etat2{
		t.Fatalf("attendu pas ce qu'on a eu")
	}

	
}
func Test_enlever_arrive(t *testing.T) {
    Liste_pays = []string{"FR","GA","GB"}

    etat := Etat{
        Visites: "101",
        Arrive: "FR",
    }

    obtenu := Enlever(etat, "GB")

    attendu := Etat{
        Visites: "100",
        Arrive: "GB",
    }

    if obtenu != attendu {
        t.Fatalf("attendu %+v, obtenu %+v", attendu, obtenu)
    }
}
func Test_Cout_dynamique(t *testing.T){

	Memo_cout = make(map[Etat]float64)
	Memo_pred = make(map[Etat]string)
	etat := Etat{
        Visites: "000",
        Arrive:  "FR",
    }

	donnees2 := map[string]donnees.Coordonnees {
		"FR": {Latitude : 46.227638, Longitude: 2.213749},
		"GA": {Latitude :-0.803689, Longitude:  11.609444},
		"GB": {Latitude : 55.378051, Longitude: -3.435973},
	}
	Liste_pays = Makelist(donnees2)

	res:= Cout_dynamique(etat, "GA")

	res_attendu := glouton.Distance(donnees.Monde["GA"], donnees.Monde["FR"])

	if res != res_attendu {
		t.Fatalf("attendu pas ce qu'on a eu")
	}


}

func Test_Pred(t *testing.T) {

    Memo_cout = make(map[Etat]float64)
    Memo_pred = make(map[Etat]string)

    Liste_pays = []string{"FR","GA","GB"}

    etat := Etat{
        Visites: "011",
        Arrive: "FR",
    }

    Cout_dynamique(etat, "FR")

    pred := Memo_pred[etat]

    if pred != "GA" {
		t.Fatalf("attendu GA, obtenu %s", pred)
	}
}