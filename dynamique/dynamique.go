package dynamique

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"math"
	"strings"
)

var Memo_cout map[Etat]float64
var Memo_pred map[Etat]string
var Liste_villes []string


type Etat struct {
	Pas_Visitees string
	Arrivee string
}



func longueur(etat Etat) int {
    return strings.Count(etat.Pas_Visitees, "1")
}

func villeavisitees(etat Etat) ([]string){
	var new []string
	for index := range etat.Pas_Visitees{
		if etat.Pas_Visitees[index] == '1' {
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
			new = new + string(etat.Pas_Visitees[j])
		}else {
			new = new + "0"
		}
	}

	new_etat := Etat{ Pas_Visitees: new, Arrivee: ville}
	return new_etat
}







func Cout_dynamique( etat Etat, depart string, Points map[string]donnees.Coordonnees) (float64){
    if longueur(etat) == 0 {

        res := glouton.Distance( Points[depart],  Points[etat.Arrivee])
		Memo_cout[etat] = res
		Memo_pred[etat] = etat.Arrivee
        return res
	}else {
        min_cout := math.Inf(1)
		ville_min := ""
		villes_a_visitees := villeavisitees(etat)

        for index := range villes_a_visitees {
			ville := villes_a_visitees[index]
			etat1 :=  Enlever(etat, ville)
			_, exists := Memo_cout[etat1]
			var res1 float64
			if exists {
			 res1 = Memo_cout[etat1] + glouton.Distance( Points[ville], Points[etat.Arrivee])
			
			}else {
				res1 = Cout_dynamique( etat1, depart, Points) + glouton.Distance( Points[ville], Points[etat.Arrivee])

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

func TSP_dynamique (depart string, dico_villes map[string]donnees.Coordonnees)([]string){

	Memo_pred = make(map[Etat]string)
	Memo_cout = make(map[Etat]float64)

	Liste_villes = donnees.Makelist(dico_villes)
	for i, v := range Liste_villes {
    if v == depart {
        Liste_villes = append(Liste_villes[:i], Liste_villes[i+1:]...)
        break
    }
}

	liste := strings.Repeat("1", len(Liste_villes))
	etat := Etat{ Pas_Visitees : liste , Arrivee : depart }

	Cout_dynamique(etat,depart, dico_villes)

	var chemin_min []string

	chemin_min = append(chemin_min, depart)
for i := 0; i < len(Liste_villes); i++ {

		ville := Memo_pred[etat]
		chemin_min = append(chemin_min, ville)
		etat = Enlever(etat, ville)
	}
chemin_min = append(chemin_min, depart)
	return chemin_min
}