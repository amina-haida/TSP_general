package temps

import (
	"TSP_general/branchandbound"
	"TSP_general/christofides"
	"TSP_general/donnees"
	"TSP_general/dynamique"
	"TSP_general/exhaustive"
	"TSP_general/glouton"
	"TSP_general/kruskal"
	"TSP_general/opt_glouton"
	"TSP_general/recuit_simule"
	"TSP_general/strip"
	"fmt"
	"time"
)

type Algo struct {
    Nom  string
    fonction func(string, map[string]donnees.Coordonnees) []string
}

func Comparaison_exacts( n int){

algos := []Algo{

    {"exhaustive", exhaustive.Exhaustive},
    {"dynamique", dynamique.TSP_dynamique},
    {"branchandbound", branchandbound.F_branchandbound},
}

depart := "FR"


temps_ville := make(map[string][]time.Duration)

if n<=11{
for i:=0; i<5 ; i++{

var dicoVilles map[string]donnees.Coordonnees

for {
    dicoVilles = donnees.Ville_aleatoire(n)
    if _, ok := dicoVilles["FR"]; ok {
        break
    }
}
for _ , algo := range algos {

		start := time.Now()
		_ = algo.fonction(depart, dicoVilles)
		duree := time.Since(start)
		temps_ville[algo.Nom] = append(temps_ville[algo.Nom],duree)

	}}

moyenne := make(map[string]time.Duration) 

for nom, liste := range temps_ville{
	moyenne[nom] = donnees.Somme(liste)/5
}

for nom,temps := range moyenne {
	fmt.Println("L'algo", nom," a une moyenne de ", temps)

}}else if n <= 15{

	for i:=0; i<5 ; i++{

var dicoVilles map[string]donnees.Coordonnees

for {
    dicoVilles = donnees.Ville_aleatoire(n)
    if _, ok := dicoVilles["FR"]; ok {
        break
    }
}
for i, algo := range algos {
		if i !=0{
		start := time.Now()
		_ = algo.fonction(depart, dicoVilles)
		duree := time.Since(start)
		temps_ville[algo.Nom] = append(temps_ville[algo.Nom],duree)

	}}}

moyenne := make(map[string]time.Duration) 

for nom, liste := range temps_ville{
	moyenne[nom] = donnees.Somme(liste)/5
}

for nom,temps := range moyenne {
	fmt.Println("L'algo", nom," a une moyenne de ", temps)
}

	}else {
	fmt.Println("Trop de villes en paramètres")
}


}

func Comparaison_heuristiques(n int) {

	algos_2 := []Algo{
		{"glouton", glouton.Fonc_glouton},
		{"opt_glouton", opt_glouton.Fonc_optGlouton},
		{"recuit_simule", recuit_simule.Recuit},
		{"strip", strip.MeilleurStrip},
		{"kruskal", kruskal.Kruskal_TSP},
		{"christofides", christofides.Christofides},
	}

	depart := "FR"

	temps_ville := make(map[string][]time.Duration)
	cout_ville := make(map[string][]float64)
	if n <= 11{

	for i := 0; i < 5; i++ {
var dicoVilles map[string]donnees.Coordonnees

for {
    dicoVilles = donnees.Ville_aleatoire(n)
    if _, ok := dicoVilles["FR"]; ok {
        break
    }
}

	
		for _, algo := range algos_2 {


			start := time.Now()
			chemin := algo.fonction(depart, dicoVilles)
			duree := time.Since(start)

			temps_ville[algo.Nom] = append(temps_ville[algo.Nom], duree)

			cout := glouton.Cout(chemin, dicoVilles)
			cout_ville[algo.Nom] = append(cout_ville[algo.Nom], cout)

		}
	}


	fmt.Println("===== Temps moyens =====")
	for nom, liste := range temps_ville {
		moyenne := donnees.Somme(liste) / time.Duration(len(liste))
		fmt.Println(nom, ":", moyenne)
	}

	
	fmt.Println("\n===== Coûts moyens =====")
	for nom, liste := range cout_ville {
		moyenne := donnees.Somme2(liste) / float64(len(liste))
		fmt.Println(nom, ":", moyenne)
	}		}else{

				for i := 0; i < 5; i++ {
var dicoVilles map[string]donnees.Coordonnees

for {
    dicoVilles = donnees.Ville_aleatoire(n)
    if _, ok := dicoVilles["FR"]; ok {
        break
    }
}

	
	
		for _, algo := range algos_2 {

			start := time.Now()
			chemin := algo.fonction(depart, dicoVilles)
			duree := time.Since(start)

			temps_ville[algo.Nom] = append(temps_ville[algo.Nom], duree)

			cout := glouton.Cout(chemin, dicoVilles)
			cout_ville[algo.Nom] = append(cout_ville[algo.Nom], cout)


		}
	}


	fmt.Println("===== Temps moyens =====")
	for nom, liste := range temps_ville {
		moyenne := donnees.Somme(liste) / time.Duration(len(liste))
		fmt.Println(nom, ":", moyenne)
	}

	
	fmt.Println("\n===== Coûts moyens =====")
	for nom, liste := range cout_ville {
		moyenne := donnees.Somme2(liste) / float64(len(liste))
		fmt.Println(nom, ":", moyenne)
	}

		}
}



func Garantie(n int){
	if n <=20{
var dicoVilles map[string]donnees.Coordonnees

for {
    dicoVilles = donnees.Ville_aleatoire(n)
    if _, ok := dicoVilles["FR"]; ok {
        break
    }
}

cout_exact := glouton.Cout(dynamique.TSP_dynamique("FR", dicoVilles), dicoVilles)

cout_k := glouton.Cout(kruskal.Kruskal_TSP("FR",dicoVilles),dicoVilles)
cout_c := glouton.Cout(christofides.Christofides("FR",dicoVilles),dicoVilles)

if cout_k <= cout_exact*2{
fmt.Println("Le cout exact est", cout_exact, "celui de kurskal est", cout_k, "il est bien <= 2*cout_opt")
}else {
	fmt.Println("Le cout exact est", cout_exact, "celui de kurskal est", cout_k, "il n'est pas <= 2*cout_opt")
}

if cout_c <= cout_exact*3/2{
fmt.Println("Le cout exact est", cout_exact, "celui de christofides est", cout_c, "il est bien <= 3/2*cout_opt")
}else {
	fmt.Println("Le cout exact est", cout_exact, "celui de christofides est", cout_c, "il n'est pas <= 3/2*cout_opt")
}
	}else{
		fmt.Println("Trop de pays en paramètres")
	}

}