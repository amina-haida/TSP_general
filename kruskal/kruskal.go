package kruskal


import(
	"TSP_general/glouton"
	"TSP_general/donnees"
	"TSP_general/dynamique"

)

type  Arrete struct{

	Poids float64
	Depart string
	Arrivee string

}


func Poids_liste(donnees map[string]donnees.Coordonnees)([]Arrete){

	Liste_pays := dynamique.Makelist(donnees ) 

	var Liste_Arrete []Arrete

	for i := range Liste_pays {
		for j:= i+1 ; j< len(Liste_pays); j++{
			Liste_Arrete = append(Liste_Arrete, Arrete{Poids : glouton.Distance(donnees[Liste_pays[i]], 
				donnees[Liste_pays[j]]), Depart : Liste_pays[i], Arrivee : Liste_pays[j] } )
		}
	}

	return Liste_Arrete
}


func Combiner(L1,L2 []Arrete)[]Arrete{

	i1:=0
	i2:=0
	var L []Arrete

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

func Tri_Arrete(Liste_Arrete []Arrete) []Arrete{

	if len(Liste_Arrete) <2 {
		return Liste_Arrete
	}else {
		milieu := len(Liste_Arrete)/2
		Liste1 := Liste_Arrete[:milieu]
		Liste2 := Liste_Arrete[milieu:]

		L1_triee := Tri_Arrete(Liste1)
		L2_triee := Tri_Arrete(Liste2)

		L_finale := Combiner(L1_triee, L2_triee)

		return L_finale

	}
}



func Kruskal_privee( donnees map[string]donnees.Coordonnees) ([]Arrete){

	Arrete_triee := Tri_Arrete(Poids_liste(donnees))
	var Arbre_liste []Arrete
	R := make(map[string]string)


	for pays := range donnees{
		R[pays] = pays
	}

	for i := range Arrete_triee{
		if R[Arrete_triee[i].Depart] != R[Arrete_triee[i].Arrivee] {

			Arbre_liste = append(Arbre_liste, Arrete_triee[i])
			stock := R[Arrete_triee[i].Arrivee]
			for pays := range donnees {
				if R[pays] == stock{
					R[pays] = Arrete_triee[i].Depart
				}
			}

		}

	}
		return Arbre_liste

}


func Conversion_dico(Arbre_liste []Arrete)map[string][]string{
	Arbre_dico := make(map[string][]string)
	for i := range Arbre_liste{
		Arrete:= Arbre_liste[i]
		Arbre_dico[Arrete.Depart] = append( Arbre_dico[Arrete.Depart],Arrete.Arrivee)
    	Arbre_dico[Arrete.Arrivee] = append( Arbre_dico[Arrete.Arrivee],Arrete.Depart)
	}

	return Arbre_dico
}

func Parcours(Arbre_dico map[string][]string, Chemin []string, sommet_actuel string, visites map[string]bool) []string{
	Chemin = append(Chemin, sommet_actuel)
	visites[sommet_actuel]=true
	for i := range Arbre_dico[sommet_actuel]{
		voisin := Arbre_dico[sommet_actuel][i]
		if visites[voisin] == false {
			Chemin = append(Chemin, voisin)
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
	return Chemin
}