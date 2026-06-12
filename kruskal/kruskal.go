package kruskal


import(
	"TSP_general/glouton"
	"TSP_general/donnees"
	"TSP_general/dynamique"
	"slices"
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
	var Chemin []Arrete
	R := make(map[string]string)


	for pays := range donnees{
		R[pays] = pays
	}

	for i := range Arrete_triee{
		if R[Arrete_triee[i].Depart] != R[Arrete_triee[i].Arrivee] {

			Chemin = append(Chemin, Arrete_triee[i])
			stock := R[Arrete_triee[i].Arrivee]
			for pays := range donnees {
				if R[pays] == stock{
					R[pays] = Arrete_triee[i].Depart
				}
			}

		}

	}
		return Chemin

}

func Kruskal(donnees map[string]donnees.Coordonnees )[]string{

	Chemin := Kruskal_privee(donnees)
	Chemin_pays := make([]string,0)

	var pays_depart string 
	var pays_arrivee string
	compteur := 0
	

	for i := 0; i<len(Chemin);i++{
		verif1:=false
		verif2:=false

		for i1:= 0; i1<len(Chemin);i1 ++{
			if i!=i1 && Chemin[i1].Depart==Chemin[i].Depart {
				verif1 = true 
			}
		}
		for i2:= 0; i2<len(Chemin);i2 ++{
			if i!=i2 && Chemin[i2].Arrivee==Chemin[i].Depart {
				verif1 = true 
			}
		}

		for i3:= 0; i3<len(Chemin);i3 ++{
			if i!=i3 && Chemin[i3].Depart==Chemin[i].Arrivee {
				verif2 = true 
			}
		}
		for i4:= 0; i4<len(Chemin);i4 ++{
			if i!=i4 && Chemin[i4].Arrivee==Chemin[i].Arrivee {
				verif2 = true 
			}
		}

		if verif1 && compteur ==0{
			pays_depart = Chemin[i].Depart
			compteur =1
		}else if verif2 && compteur ==0{
			pays_depart = Chemin[i].Arrivee
			compteur =1
		}
		if verif1 && compteur ==1{
			pays_arrivee = Chemin[i].Depart
		}else if verif2 && compteur ==1{
			pays_arrivee = Chemin[i].Arrivee
		}
	}
		Chemin_pays= append(Chemin_pays, pays_depart)
		for j:= 0; j<len(donnees) ; j++{
			if Chemin[j].Depart == Chemin_pays[len(Chemin_pays)-1] {
				Chemin_pays = append(Chemin_pays, Chemin[j].Arrivee)
				Chemin = slices.Delete(Chemin, j, j+1)
			}else if Chemin[j].Arrivee == Chemin_pays[len(Chemin_pays)-1]{
				Chemin_pays = append(Chemin_pays, Chemin[j].Depart)
				Chemin = slices.Delete(Chemin, j, j+1)			
			}
		}
	Chemin_pays=append(Chemin_pays, pays_arrivee)

	return Chemin_pays
}