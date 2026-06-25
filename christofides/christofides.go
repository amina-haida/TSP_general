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

	couplages := make([]kruskal.Arete,0)
	for len(S) > 1 {
		var sommet_i string
		for k := range S {
			sommet_i = k
			break
		}

    sommet_f := glouton.PlusProche(sommet_i, S, dicoVilles)

    couplages = append(couplages, kruskal.Arete{
		        Poids: glouton.Distance(
            dicoVilles[sommet_i],
            dicoVilles[sommet_f],
        ),
        Depart: sommet_i,
        Arrivee: sommet_f,

    })

    delete(S, sommet_i)
    delete(S, sommet_f)
}
		

return couplages
}

func union_aretes(couplages []kruskal.Arete, MST []kruskal.Arete)[]kruskal.Arete{
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
	cycle := []string{depart}

	for {

		// sécurité
		if voisins, ok := graphe[sommet]; !ok || len(voisins) == 0 {
			break
		}

		voisin := graphe[sommet][0]

		cycle = append(cycle, voisin)

		graphe[sommet] = delete_liste(graphe[sommet], voisin)

		if len(graphe[sommet]) == 0 {
			delete(graphe, sommet)
		}

		graphe[voisin] = delete_liste(graphe[voisin], sommet)

		if len(graphe[voisin]) == 0 {
			delete(graphe, voisin)
		}

		sommet = voisin

		if sommet == depart {
			break
		}
	}

	return cycle, graphe
}


func combiner(cycle1 []string, cycle2 []string, sommet string)[]string{
	occurrence := 0
	new_cycle := make([]string, 0)
	for _,sommet2 := range cycle1{
		if sommet == sommet2 && occurrence==0{
			occurrence ++
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
			break
		}
	}
	return retour 
}
func cycle_eulérien(depart string, graphe map[string][]string)[]string{

	cycle, graphe := Cycle(depart, graphe)

	for len(graphe)>0 {

		sommet_suivant := suivant(cycle, graphe)

		if sommet_suivant == "" {
			break
		}

		new_cycle, nouveauGraphe := Cycle(sommet_suivant, graphe)

		cycle = combiner(cycle, new_cycle, sommet_suivant)

		graphe = nouveauGraphe
	}

	return cycle
}

func Cycle_hamiltonien(cycle_eul []string)[]string{
	cycle := make([]string, 0)
	visites := make(map[string]bool)
	for _,sommet := range cycle_eul{
		if visites[sommet]!= true {
			cycle = append(cycle, sommet)
			visites[sommet ]= true
		}

	}
	return cycle
}

func Christofides(depart string, dicoVilles map[string]donnees.Coordonnees)[]string{

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

	cycle_hamiltonien := Cycle_hamiltonien(cycle_eul)

	cycle_hamiltonien = append(cycle_hamiltonien, depart)
		
	return cycle_hamiltonien
	}
