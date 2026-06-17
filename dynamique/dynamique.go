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
var Liste_pays []string


type Etat struct {
	Visites string
	Arrive string
}

func Makelist( donnees map[string]donnees.Coordonnees) ([]string){
	var new []string
	for pays := range donnees {
		new = append(new, pays)
	}
	slices.Sort(new)
	return new
}


func longueur(etat Etat) int {
    return strings.Count(etat.Visites, "1")
}

func paysvisites(etat Etat) ([]string){
	var new []string
	for index := range etat.Visites{
		if etat.Visites[index] == '1' {
			new = append(new, Liste_pays[index])

		}
	}

		return new
	}

func Enlever(etat Etat, pays string)(Etat){
	index := 0
	for i := range Liste_pays {
		if Liste_pays[i] == pays {
			index = i
		}
	}
	var new string
	for j := range Liste_pays{
		if j != index {
			new = new + string(etat.Visites[j])
		}else {
			new = new + "0"
		}
	}

	new_etat := Etat{ Visites: new, Arrive: pays}
	return new_etat
}







func Cout_dynamique( etat Etat ,depart string) (float64){
    if longueur(etat) == 0 {
        res := glouton.Distance( donnees.Monde[depart],  donnees.Monde[etat.Arrive])
		Memo_cout[etat] = res
        return res
	}else {
        min_cout := math.Inf(1)
		pays_min := ""
		pays_visites := paysvisites(etat)

        for index := range pays_visites {
			pays := pays_visites[index]
			etat1 :=  Enlever(etat, pays)
			_, exists := Memo_cout[etat1]
			var res1 float64
			if exists {
			 res1 = Memo_cout[etat1] + glouton.Distance( donnees.Monde[pays],donnees.Monde[etat.Arrive])
			
			}else {
				res1 = Cout_dynamique( etat1, depart) + glouton.Distance( donnees.Monde[pays],donnees.Monde[etat.Arrive])

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

	Liste_pays = Makelist(map_pays)

	liste := strings.Repeat("1", len(Liste_pays))
	etat := Etat{ Visites : liste , Arrive : depart }

	Cout_dynamique(etat,"FR")

	var chemin_min []string
	chemin_min = append(chemin_min, depart)

	for i :=0; i < len(Makelist(map_pays)); i++ {

		pays := Memo_pred[etat]
		chemin_min = append(chemin_min, pays)
		etat = Enlever(etat, pays)
}
		chemin_min = append(chemin_min, "FR")
		return chemin_min
}