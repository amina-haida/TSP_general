package troncon

import "fmt"

type Troncon struct {
	Depart  string
	Arrivee string
}

var Reseau_ferre []Troncon


func Troncons_rec(precedente string, courante string) {
	fmt.Println(courante)

	for _, troncon := range Reseau_ferre {
		u := troncon.Depart
		v := troncon.Arrivee
		if u == courante && v != precedente {
			Troncons_rec(u, v)
		}
	}
}

func Troncons_rec2(précédente string, courante string, chemin *[]string) {
	*chemin = append(*chemin, courante)
	for _, troncon := range Reseau_ferre {
		u := troncon.Depart
		v := troncon.Arrivee
		if u == courante && v != précédente {
			Troncons_rec2(u, v, chemin)
		}
	}

}

func Chercher_ville(precedente string, courante string, destination string, chemin *[]string) {

	if courante == destination {
		*chemin = append(*chemin, courante)
		return

	}

	for _, troncon := range Reseau_ferre {
		u := troncon.Depart
		v := troncon.Arrivee
		if u == courante && v != precedente {
			Chercher_ville(u, v, destination, chemin)
			if len(*chemin) != 0 {
				*chemin = append(*chemin, courante)
				return
			}
		}
	}
}
