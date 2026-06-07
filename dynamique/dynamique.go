package dynamique

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"math"
	"strings"
	"slices"
)

var memo_cout map[Etat]float64
var memo_pred map[Etat]string


type Etat struct {
	visites string
	arrive string
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
    return strings.Count(etat.visites, "1")
}

func pays(etat Etat, liste_pays []string) ([]string){
	var new []string
	for index := range etat.visites{
		if etat.visites[index] == '1' {
			new = append(new, liste_pays[index])

		}
	}

		return new
	}

func enlever(etat Etat, pays string, liste_pays []string)(Etat){
	index := 0
	for i := range liste_pays {
		if liste_pays[i] == pays {
			index = i
		}
	}
	var new string
	for j := range liste_pays{
		if j != index {
			new = new + string(etat.visites[j])
		}else {
			new = new + "0"
		}
	}

	new_etat := Etat{ visites: new, arrive: pays}
	return new_etat
}







func Cout_dynamique( etat Etat ,depart string, liste_pays []string) (float64){
    if longueur(etat) == 0 {
        res := glouton.Distance(depart, etat.arrive, donnees.Pays)
		memo_cout[etat] = res
        return res
	}else {
        min_cout := math.Inf(1)
		pays_min := "FR"
		pays_visites := pays(etat, liste_pays)

        for index := range len(pays_visites) {
			pays := pays_visites[index]
			etat1 :=  enlever(etat, pays, liste_pays)
			_, exists := memo_cout[etat1]
			var res1 float64
			if exists {
			 res1 = memo_cout[etat1] + glouton.Distance( pays, etat.arrive, donnees.Pays)
			
			}else {
				res1 = Cout_dynamique( etat1, depart, liste_pays) + glouton.Distance( pays,etat.arrive, donnees.Pays)

			}
            if res1 < min_cout {
				min_cout = res1 
                pays_min = pays 
			}

			} 
            memo_cout[etat] = min_cout 
			memo_pred[etat] = pays_min
            return min_cout
            
        
	}

}