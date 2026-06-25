package christofides

import ("TSP_general/donnees"
"TSP_general/kruskal"
"TSP_general/glouton")

func delete_liste(liste []string, sommet string)[]string{
	new_liste := make([]string,0)
	for _,ville := range liste {
		if ville != sommet {
			new_liste = append(new_liste, ville)
		}
	}
	return new_liste
}

func couplages(S map[string]donnees.Coordonnees,  dicoVilles map[string]donnees.Coordonnees) []kruskal.Arete{
	n := len(S)
	couplages := make([]kruskal.Arete,0)
	for len(couplages) < n/2{
		for sommet_i := range S{
			sommet_f := glouton.PlusProche(sommet_i,S, dicoVilles)
			delete(S, "sommet_i")
			delete(S, "sommet_f")
			couplages = append(couplages, kruskal.Arete{Poids : glouton.Distance(dicoVilles[sommet_f], dicoVilles[sommet_i]),Depart: sommet_i, Arrivee: sommet_f})
			}
		}

return couplages
}

func union_aretes( couplages []kruskal.Arete, MST []kruskal.Arete)[]kruskal.Arete{
			arete_multigraphe := make([]kruskal.Arete,0)


		for _,arete := range MST {
			arete_multigraphe = append(arete_multigraphe,arete )

		}

		for _,arete := range couplages {
			arete_multigraphe = append(arete_multigraphe,arete )
		}
return arete_multigraphe
}

func Cycle(depart string, graphe map[string][]string)([]string,map[string][]string){
	sommet := depart 
	cycle := []string {depart}
	verif := true 
	for verif {
		voisin := graphe[sommet][0]
		cycle = append(cycle, voisin)
		graphe[sommet] = delete_liste(graphe[sommet], voisin)
		if len(graphe[sommet]) == 0{
			delete(graphe, sommet)
		}
		graphe[voisin] = delete_liste(graphe[voisin], sommet)
		if len(graphe[voisin]) == 0{
			delete(graphe, voisin)
		}
		sommet = voisin 
		if voisin == depart {
			verif = false
		}
	}
	return cycle, graphe
}


func combiner(cycle1 []string, cycle2 []string, sommet string)[]string{
	occurrence := 0
	new_cycle := make([]string, 0)
	for _,sommet2 := range cycle1{
		if sommet == sommet2 && occurrence==0{
			for _,sommet3 := range cycle2{
				new_cycle = append(new_cycle, sommet3)
			}
		}else {
			new_cycle = append(new_cycle, sommet2)
		}
	}
	return new_cycle
}
func suivant(cycle []string, graphe map[string][]string)string{
	retour := ""
	for _,sommet := range cycle {
		if len(graphe[sommet])>0{
			retour = sommet
		}
	}
	return retour 
}
func cycle_eulérien(depart string, graphe map[string][]string)[]string{
	var cycle []string
	cycle,graphe = Cycle(depart, graphe)
	for len(graphe)>0 {
		sommet_suivant := suivant(cycle, graphe) 
		var new_cycle []string 
		new_cycle, graphe = Cycle(sommet_suivant, graphe)
		cycle = combiner(cycle, new_cycle, sommet_suivant)


}
return cycle }

func christofides(depart string, dicoVilles map[string]donnees.Coordonnees){

	arbre := kruskal.Kruskal_privee(dicoVilles)

	Arbre := kruskal.Conversion_dico(arbre)

	S := make(map[string]donnees.Coordonnees)

	for sommet,voisins := range Arbre {
		if len(voisins)%2 != 0 {
			S[sommet] = dicoVilles[sommet]
		}
	}
	graphe_couplages := couplages(S, dicoVilles)

	arete_multigraphe := union_aretes(graphe_couplages, arbre)

	multigraphe := kruskal.Conversion_dico(arete_multigraphe)

    cycle_eul := cycle_eulérien(depart, multigraphe)
		
	}
