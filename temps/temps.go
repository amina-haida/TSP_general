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
	"math"
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

dicoVilles := donnees.Ville_aléatoire(n)
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

}}else if n <=20 {
	dicoVilles := donnees.Ville_aléatoire(n)
for _ , algo := range algos {

		start := time.Now()
		_ = algo.fonction(depart, dicoVilles)
		duree := time.Since(start)
		temps_ville[algo.Nom] = append(temps_ville[algo.Nom],duree)

	}

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
	erreurs := make(map[string][]float64)

	for i := 0; i < 5; i++ {

		dicoVilles := donnees.Ville_aléatoire(n)


		coutExact := glouton.Cout(dynamique.TSP_dynamique(depart, dicoVilles))
	
		for _, algo := range algos_2 {

			start := time.Now()
			chemin := algo.fonction(depart, dicoVilles)
			duree := time.Since(start)

			temps_ville[algo.Nom] = append(temps_ville[algo.Nom], duree)

			cout := glouton.Cout(chemin, dicoVilles)
			cout_ville[algo.Nom] = append(cout_ville[algo.Nom], cout)

			erreur := math.Abs(cout-coutExact) / coutExact
			erreurs[algo.Nom] = append(erreurs[algo.Nom], erreur)
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

	
	fmt.Println("\n===== Erreur moyenne par rapport à l'optimum =====")
	for nom, liste := range erreurs {
		moyenne := donnees.Somme2(liste) / float64(len(liste))
		fmt.Printf("%s : %.2f %%\n", nom, 100*moyenne)
	}
}