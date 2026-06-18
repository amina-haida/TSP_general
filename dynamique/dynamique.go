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







func Cout_dynamique( etat Etat, ville_suivante string) (float64){
    if longueur(etat) == 0 {

        res := glouton.Distance( donnees.Monde[ville_suivante],  donnees.Monde[etat.Arrivee])
		Memo_cout[etat] = res
        return res
	}else {
        min_cout := math.Inf(1)
		ville_min := ""
		villes_visitees := villevisitees(etat)

        for index := range villes_visitees {
			ville := villes_visitees[index]
			etat1 :=  Enlever(etat, ville)
			_, exists := Memo_cout[etat1]
			var res1 float64
			if exists {
			 res1 = Memo_cout[etat1] + glouton.Distance( donnees.Monde[ville],donnees.Monde[etat.Arrivee])
			
			}else {
				res1 = Cout_dynamique( etat1, ville_suivante) + glouton.Distance( donnees.Monde[ville],donnees.Monde[etat.Arrivee])

			}
            if res1 < min_cout {
				min_cout = res1 
                ville_min = ville 
			}

			} 
            Memo_cout[etat] = min_cout 
			Memo_pred[etat] = ville_min
            return min_cout
            
        
	}

}

func TSP_dynamique (dico_villes map[string]donnees.Coordonnees, depart string)([]string){

	Memo_pred = make(map[Etat]string)
	Memo_cout = make(map[Etat]float64)

	Liste_villes = Makelist(dico_villes)

	liste := strings.Repeat("1", len(Liste_villes))
	etat := Etat{ Visitees : liste , Arrivee : depart }

	Cout_dynamique(etat,depart)

	var chemin_min []string
	chemin_min = append(chemin_min, depart)

	for i :=0; i < len(Makelist(dico_villes)); i++ {

		ville := Memo_pred[etat]
		chemin_min = append(chemin_min, ville)
		etat = Enlever(etat, ville)
}

		return chemin_min
}