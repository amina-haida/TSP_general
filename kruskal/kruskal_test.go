package kruskal

import ("testing"
"TSP_general/donnees"

)

func Test_Poids_list(t *testing.T){
			map_villes := make(map[string]donnees.Coordonnees, 0)
	map_villes["FR"] = donnees.Europe()["FR"]
	map_villes["ES"] = donnees.Europe()["ES"]
	map_villes["UK"] = donnees.Europe()["UK"]
	map_villes["IT"] = donnees.Europe()["IT"]
	map_villes["LI"] = donnees.Europe()["LI"]
	map_villes["MT"] = donnees.Europe()["MT"]
	map_villes["NL"] = donnees.Europe()["NL"]
	map_villes["NO"] = donnees.Europe()["NO"]
	map_villes["PL"] = donnees.Europe()["PL"]
	map_villes["PT"] = donnees.Europe()["PT"]
	map_villes["SE"] = donnees.Europe()["SE"]
	res := Poids_liste(map_villes)
	n := len(map_villes)
	if len(res) != n * (n-1)/2 {
		t.Fatal("résultat incorrect")
	}

}

func Test_Combiner(t *testing.T){
villes1 := []Arete{
	{Poids: 2, Depart: "a", Arrivee: "b"},
		{Poids: 2, Depart: "b", Arrivee: "c"},
			{Poids: 2, Depart: "a", Arrivee: "c"},
}	
villes2 := []Arete{
	{Poids: 2, Depart: "a", Arrivee: "d"},
		{Poids: 2, Depart: "d", Arrivee: "e"},
			{Poids: 2, Depart: "a", Arrivee: "e"},
}	

res := Combiner(villes1,villes2)

if len(res ) != 6 {
		t.Fatal("résultat incorrect")
}
}

func Test_dico_conv(t *testing.T){
	villes1 := []Arete{
	{Poids: 2, Depart: "a", Arrivee: "b"},
		{Poids: 2, Depart: "b", Arrivee: "c"},
			{Poids: 2, Depart: "a", Arrivee: "c"},
}

res := Conversion_dico(villes1)

if len(res) != 3 {
	t.Fatal("résultat incorrect")

}

if len(res["a"]) != 2 {
		t.Fatal("résultat incorrect")
}


}

func Test_Parcours(t *testing.T){
		villes1 := []Arete{
	{Poids: 2, Depart: "a", Arrivee: "b"},
		{Poids: 2, Depart: "b", Arrivee: "c"},
			{Poids: 2, Depart: "a", Arrivee: "c"},
}

graphe := Conversion_dico(villes1)
visites :=make(map[string]bool)
chemin := []string{}
res := Parcours(graphe, chemin, "a", visites)

if len(res) != 3{
	t.Fatal("résultat incorrect")
}

}