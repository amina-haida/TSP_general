package kruskal


import(
	"TSP_general/glouton"
	"TSP_general/donnees"
	"slices"

)

type  Arete struct{

	Poids float64
	Depart string
	Arrivee string

}


func Poids_liste(dico_villes map[string]donnees.Coordonnees)([]Arete){

	Liste_pays := donnees.Makelist(dico_villes ) 

	var Liste_Arrete []Arete

	for i := range Liste_pays {
		for j:= i+1 ; j< len(Liste_pays); j++{
			Liste_Arrete = append(Liste_Arrete, Arete{Poids : glouton.Distance(dico_villes[Liste_pays[i]], 
				dico_villes[Liste_pays[j]]), Depart : Liste_pays[i], Arrivee : Liste_pays[j] } )
		}
	}

	return Liste_Arrete
}


func Combiner(L1,L2 []Arete)[]Arete{

	i1:=0
	i2:=0
	var L []Arete

	for i:=0 ; i<len(L1)+len(L2) ; i++{

		if i1 == len(L1) {
			L=append(L, L2[i2])
			i2++

		} else if i2==len(L2){
			L=append(L, L1[i1])

			i1++
		}else if L1[i1].Poids < L2[i2].Poids {

			L=append(L,L1[i1] )

			i1++
		} else {

			L = append(L, L2[i2])
			i2++
		}
	}
	return L
}
// On ne prouve pas Tri_Arete comme cette fonction correspond au tri fusion 
func Tri_Arete(Liste_Arete []Arete) []Arete{

	if len(Liste_Arete) <2 {
		return Liste_Arete
	}else {
		milieu := len(Liste_Arete)/2
		Liste1 := Liste_Arete[:milieu]
		Liste2 := Liste_Arete[milieu:]

		L1_triee := Tri_Arete(Liste1)
		L2_triee := Tri_Arete(Liste2)

		L_finale := Combiner(L1_triee, L2_triee)

		return L_finale

	}
}



func Kruskal_privee( donnees map[string]donnees.Coordonnees) ([]Arete){

	Arete_triee := Tri_Arete(Poids_liste(donnees))
	var Arbre_liste []Arete
	R := make(map[string]string)


	for pays := range donnees{
		R[pays] = pays
	}

	for i := range Arete_triee{
		if R[Arete_triee[i].Depart] != R[Arete_triee[i].Arrivee] {

			Arbre_liste = append(Arbre_liste, Arete_triee[i])
			stock := R[Arete_triee[i].Arrivee]
			for pays := range donnees {
				if R[pays] == stock{
					rep := R[Arete_triee[i].Depart]
stock := R[Arete_triee[i].Arrivee]

for pays := range donnees {
    if R[pays] == stock {
        R[pays] = rep
    }
}
				}
			}

		}

	}
		return Arbre_liste

}

func Conversion_dico(Arbre_liste []Arete) map[string][]string {
    Arbre_dico := make(map[string][]string)

    for _, arete := range Arbre_liste {
        Arbre_dico[arete.Depart] = append(Arbre_dico[arete.Depart], arete.Arrivee)
        Arbre_dico[arete.Arrivee] = append(Arbre_dico[arete.Arrivee], arete.Depart)
    }

    for sommet := range Arbre_dico {
        slices.Sort(Arbre_dico[sommet])
    }

    return Arbre_dico
}

func Parcours(Arbre_dico map[string][]string, Chemin []string, sommet_actuel string, visites map[string]bool) []string{
	Chemin = append(Chemin, sommet_actuel)
	visites[sommet_actuel]=true
	for i := range Arbre_dico[sommet_actuel]{
		voisin := Arbre_dico[sommet_actuel][i]
		if visites[voisin] != true {
			Chemin = Parcours( Arbre_dico, Chemin, voisin, visites)
		}
	}
return Chemin
}

func Kruskal_TSP(depart string, donnees map[string]donnees.Coordonnees ) []string{
	Arbre_liste := Kruskal_privee(donnees)
	Arbre_dico := Conversion_dico(Arbre_liste)
	Chemin := make([]string,0)
	visites := make(map[string]bool)
	Chemin = Parcours( Arbre_dico, Chemin, depart, visites)
	Chemin = append(Chemin, depart)
	return Chemin
}