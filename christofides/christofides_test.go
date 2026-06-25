package christofides

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"TSP_general/kruskal"
	"testing"
)

func Test_delete_liste(t *testing.T){
	l1 := []string{"a","b","c"}
	l2 := []string{"a","a","a"}
	res1 := delete_liste(l1, "a")
	res1_vrai := []string{"b","c"}
	res2 := delete_liste(l2, "a")
	res2_vrai := []string{}



	for i := range res1 {
		if res1[i] != res1_vrai[i]{
			t.Fatal("résultat incorrect")

		}
	}
	if len(res2_vrai) != len(res2){
		t.Fatal("résultat incorrect")
	}
	}


func Test_couplages(t *testing.T){

		dico_villes := make(map[string]donnees.Coordonnees, 0 )
	dico_villes["FR"] = donnees.Europe()["FR"]
	dico_villes["ES"] = donnees.Europe()["ES"]
	dico_villes["UK"] = donnees.Europe()["UK"]
	dico_villes["IT"] = donnees.Europe()["IT"]
	dico_villes["LI"] = donnees.Europe()["LI"]

	S := make(map[string]donnees.Coordonnees, 0 )
	S["FR"] = donnees.Europe()["FR"]
	S["ES"] = donnees.Europe()["ES"]
	S["UK"] = donnees.Europe()["UK"]
	S["IT"] = donnees.Europe()["IT"]

	res := couplages(S, dico_villes )

	if len(res) != 2 {
				t.Fatal("résultat incorrect")
	}

}


func Test_union_arete(t *testing.T){
			dico_villes := make(map[string]donnees.Coordonnees, 0 )
	dico_villes["FR"] = donnees.Europe()["FR"]
	dico_villes["ES"] = donnees.Europe()["ES"]
	dico_villes["UK"] = donnees.Europe()["UK"]
	dico_villes["IT"] = donnees.Europe()["IT"]
	dico_villes["LI"] = donnees.Europe()["LI"]

	graphe1 := []kruskal.Arete{
		{
		        Poids: glouton.Distance(
            dico_villes["FR"],
            dico_villes["ES"],
        ),
        Depart: "FR",
        Arrivee: "ES",
	},
	}
	graphe2 := []kruskal.Arete{
	{
		        Poids: glouton.Distance(
            dico_villes["IT"],
            dico_villes["UK"],
        ),
        Depart: "IT",
        Arrivee: "UK",
	},
	}

	res := union_aretes(graphe1, graphe2)

	res_vrai := []kruskal.Arete{

	{
		        Poids: glouton.Distance(
            dico_villes["IT"],
            dico_villes["UK"],
        ),
        Depart: "IT",
        Arrivee: "UK",
	},
		{
		        Poids: glouton.Distance(
            dico_villes["FR"],
            dico_villes["ES"],
        ),
        Depart: "FR",
        Arrivee: "ES",
	},
	}

	for i := range res{

		if res[i].Poids != res_vrai[i].Poids{
			t.Fatal("résultat incorrect")
		}else if res[i].Depart != res_vrai[i].Depart{
			t.Fatal("résultat incorrect")
		} else if  res[i].Arrivee != res_vrai[i].Arrivee{
			t.Fatal("résultat incorrect")
		}


	}
}

func Test_cycle(t *testing.T){

dico_villes := make(map[string]donnees.Coordonnees, 0 )
	dico_villes["FR"] = donnees.Europe()["FR"]
	dico_villes["ES"] = donnees.Europe()["ES"]
	dico_villes["UK"] = donnees.Europe()["UK"]
	dico_villes["IT"] = donnees.Europe()["IT"]
	dico_villes["LI"] = donnees.Europe()["LI"]
	dico := kruskal.Conversion_dico(kruskal.Poids_liste(dico_villes))
	res, _  := Cycle( "FR", dico)

	if res[0] != res[len(res)-1]{
					t.Fatal("résultat incorrect")
	}



}


func Test_Combiner(t *testing.T){
villes1 := []kruskal.Arete{
	{Poids: 2, Depart: "a", Arrivee: "b"},
		{Poids: 2, Depart: "b", Arrivee: "c"},
			{Poids: 2, Depart: "a", Arrivee: "c"},
}	
villes2 := []kruskal.Arete{
	{Poids: 2, Depart: "a", Arrivee: "d"},
		{Poids: 2, Depart: "d", Arrivee: "e"},
			{Poids: 2, Depart: "a", Arrivee: "e"},
}	
cycle1,_ := Cycle("b", (kruskal.Conversion_dico(villes1)))
cycle2,_ := Cycle("a", (kruskal.Conversion_dico(villes2)))

res := combiner(cycle1, cycle2, "a")

	if res[0] != res[len(res)-1]{
					t.Fatal("résultat incorrect")
	}

	indice :=0

	for i := range res {

		if res[i] == "a"{
			indice = i
			break
	
		}else {
			if res[i] != cycle1[i]{
				t.Fatal("résultat incorrect")
			
			}
		}
	}

	for i := indice; i< indice + len(cycle2); i++ {
	
		if res[i] != cycle2[i-indice]{
				t.Fatal("résultat incorrect")				
			}
		}
for i := indice + len(cycle2) ; i<len(res); i++{
		indice_cycle1 := i+1-len(cycle2)
	if res[i] != cycle1[indice_cycle1]{
		t.Fatal("résultat incorrect")	
	}
}

	}

func Test_cycle_hamiltonien(t *testing.T){
	cycle := []string{"a","a","b","c","b"}

	res := Cycle_hamiltonien(cycle)

	res_vrai := []string{"a","b","c"}

	for i := range res_vrai{
		if res_vrai[i] != res[i] {
			t.Fatal("résultat incorrect")
		}
	}
}
