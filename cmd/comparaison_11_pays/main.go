package main 
import (
		"TSP_general/donnees"
		"TSP_general/glouton"
		"TSP_general/opt_glouton"
		"TSP_general/recuit_simule"
		"TSP_general/kruskal"
		"TSP_general/exhaustive"
		"TSP_general/dynamique"
		"TSP_general/branchandbound"

		"fmt"
		"math"
		"time"
)

type Algo struct {
    Nom  string
    fonction func(string, map[string]donnees.Coordonnees) []string
}


func main(){



algos := []Algo{

    {"exhaustive", exhaustive.Exhaustive},
    {"dynamique", dynamique.TSP_dynamique},
    {"branchandbound", branchandbound.F_branchandbound},
}



algos_2 := []Algo{    
	{"glouton", glouton.Fonc_glouton},
	{"opt_glouton",opt_glouton.Fonc_optGlouton},

    {"recuit_simule", 	recuit_simule.Recuit},

    {"kruskal", kruskal.Kruskal_TSP},

}

depart := "FR"


temps_15_pays := make(map[string][]time.Duration)

i:= 0
for i<5{
	i++
	dicoVilles := donnees.Ville_aléatoire(15)
for _ , algo := range algos {

		start := time.Now()
		_ = algo.fonction(depart, dicoVilles)
		duree := time.Since(start)
		temps_15_pays[algo.Nom] = append(temps_15_pays[algo.Nom],duree)

	}}
i = 0
moyenne := make(map[string]time.Duration) 

for nom, liste := range temps_15_pays{
	moyenne[nom] = donnees.Somme(liste)/5
}

for nom,temps := range moyenne {
	fmt.Println("L'algo", nom," a une moyenne de ", temps)

}


temps_15_pays2 := make(map[string][]time.Duration)
cout_15_pays2 := make(map[string][]float64)

moyenne2_temps := make(map[string]time.Duration) 
moyenne2_cout := make(map[string]float64) 

liste_cout_exacts := make([]float64, 0)

exactitude := make(map[string]float64) 

for i<5{
	i++
dicoVilles := donnees.Ville_aléatoire(15)
for _, algo := range algos_2 {


		start := time.Now()
		chemin := algo.fonction(depart, dicoVilles)
		duree := time.Since(start)
		temps_15_pays2[algo.Nom] = append(temps_15_pays2[algo.Nom],duree)
		cout := glouton.Cout(chemin)
		cout_15_pays2[algo.Nom] = append(cout_15_pays2[algo.Nom], cout)
		cout_exact := glouton.Cout(branchandbound.F_branchandbound(depart, dicoVilles))
		liste_cout_exacts = append(liste_cout_exacts, cout_exact)

	
	}}



for nom, liste := range temps_15_pays2{
	moyenne2_temps[nom] = donnees.Somme(liste)/5

}

for nom,temps := range moyenne2_temps {
	fmt.Println("L'algo", nom," a une moyenne de ", temps)

}



for nom, liste := range cout_15_pays2{
	moyenne2_cout[nom] = donnees.Somme2(liste)/5
	
}

i=0
for nom, cout2 := range moyenne2_cout{
	cout_exact := liste_cout_exacts[i]
	i++
	exactitude[nom] = math.Abs(cout2-cout_exact) / cout_exact

}

}