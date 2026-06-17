package dynamique

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"math"
	"strings"
	"slices"
)

var Memo_cout map[Etat]float64
var Memo_pred map[Etat]string
var Liste_villes []string


type Etat struct {
	Visitees string
	Arrivee string
}

func Makelist( donnees map[string]donnees.Coordonnees) ([]string){
	var new []string
	for ville := range donnees {
		new = append(new, ville)
	}
	slices.Sort(new)
	return new
}


func longueur(etat Etat) int {
    return strings.Count(etat.Visitees, "1")
}

func villevisitees(etat Etat) ([]string){
	var new []string
	for index := range etat.Visitees{
		if etat.Visitees[index] == '1' {
			new = append(new, Liste_villes[index])

		}
	}

		return new
	}

func Enlever(etat Etat, ville string)(Etat){
	index := 0
	for i := range Liste_villes {
		if Liste_villes[i] == ville {
			index = i
		}
	}
	var new string
	for j := range Liste_villes{
		if j != index {
			new = new + string(etat.Visitees[j])
		}else {
			new = new + "0"
		}
	}

	new_etat := Etat{ Visitees: new, Arrivee: ville}
	return new_etat
}







func Cout_dynamique( etat Etat, depart string) (float64){
    if longueur(etat) == 0 {

        res := glouton.Distance( donnees.Monde[depart],  donnees.Monde[etat.Arrivee])
		Memo_cout[etat] = res
        return res
	}else {
        min_cout := math.Inf(1)
		pays_min := ""
		pays_visites := villevisitees(etat)

        for index := range pays_visites {
			pays := pays_visites[index]
			etat1 :=  Enlever(etat, pays)
			_, exists := Memo_cout[etat1]
			var res1 float64
			if exists {
			 res1 = Memo_cout[etat1] + glouton.Distance( donnees.Monde[pays],donnees.Monde[etat.Arrivee])
			
			}else {
				res1 = Cout_dynamique( etat1, depart) + glouton.Distance( donnees.Monde[pays],donnees.Monde[etat.Arrivee])

			}
            if res1 < min_cout {
				min_cout = res1 
                pays_min = pays 
			}

			} 
            Memo_cout[etat] = min_cout 
			Memo_pred[etat] = pays_min
            return min_cout
            
        
	}

}

func TSP_dynamique (map_pays map[string]donnees.Coordonnees, depart string)([]string){

	Memo_pred = make(map[Etat]string)
	Memo_cout = make(map[Etat]float64)

	Liste_villes = Makelist(map_pays)

	liste := strings.Repeat("1", len(Liste_villes))
	etat := Etat{ Visitees : liste , Arrivee : depart }

	Cout_dynamique(etat,"FR")

	var chemin_min []string
	chemin_min = append(chemin_min, depart)

	for i :=0; i < len(Makelist(map_pays)); i++ {

		ville := Memo_pred[etat]
		chemin_min = append(chemin_min, ville)
		etat = Enlever(etat, ville)
}
		chemin_min = append(chemin_min, "FR")
		return chemin_min
}