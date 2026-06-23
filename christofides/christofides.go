package christofides

import ("TSP_general/donnees"
"TSP_general/kruskal"
"TSP_general/glouton")

func christofides(depart string, dicoVilles map[string]donnees.Coordonnees){

	arbre := kruskal.Kruskal_privee(dicoVilles)

	Arbre := kruskal.Conversion_dico(arbre)

	non_visites := make(map[string]donnees.Coordonnees)

	for sommet,voisins := range Arbre {
		if len(voisins)%2 != 0 {
			non_visites[sommet] = dicoVilles[sommet]
		}
	}

	n := len(non_visites)
	couplages := make([]kruskal.Arrete,0)

	for len(couplages) < n/2{
		for sommet_i,_ := range non_visites{
			sommet_f := glouton.PlusProche(sommet_i,non_visites,dicoVilles)
			delete(non_visites, "sommet_i")
			delete(non_visites, "sommet_f")
			couplages = append(couplages, kruskal.Arrete{Poids : glouton.Distance(dicoVilles[sommet_f], dicoVilles[sommet_i]),Depart: sommet_i, Arrivee: sommet_f})
			}
		}

		arrete_multigraphe := make([]kruskal.Arrete,0)

		for _,arrete := range arbre {
			arrete_multigraphe = append(arrete_multigraphe,arrete )

		}

		for _,arrete := range couplages {
			arrete_multigraphe = append(arrete_multigraphe,arrete )
		}
		multigraphe := kruskal.Conversion_dico(arrete_multigraphe)

		
	}

