package main 
import (
		"TSP_general/donnees"
		"TSP_general/glouton"
		"TSP_general/opt_glouton"
		"TSP_general/kruskal"
		"TSP_general/exhaustive"
		"TSP_general/dynamique"
		"TSP_general/branchandbound"

	"fmt"
	"math"
	"time"
)

func main(){

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

	dicoVilles:=map_villes

type Algo struct {
    Nom  string
    fonction func(string, map[string]donnees.Coordonnees) []string
}

algos := []Algo{

    {"exhaustive", exhaustive.Exhaustive},
    {"dynamique", dynamique.TSP_dynamique},
    {"branchandbound", branchandbound.F_branchandbound},
}


depart := "FR"
	temps_11_pays := make([]time.Duration,0)

var chemin []string
	for _, algo := range algos {
		start := time.Now()
		chemin = algo.fonction(depart, dicoVilles)
		duree := time.Since(start)
		temps_11_pays = append(temps_11_pays,duree)

	}
	fmt.Println("Comparaison algo exacts pour 11 pays", temps_11_pays)
cout := glouton.Cout(chemin)

algos_2 := []Algo{    
	{"glouton", glouton.Fonc_glouton},

    {"kruskal", kruskal.Kruskal_TSP},

}

	temps_11_pays2 := make([]time.Duration,0)
	exactitude := make([]float64,0)
	chemins := make([][]string,0)

	for _, algo := range algos_2 {
		start := time.Now()
		chemin2 := algo.fonction(depart, dicoVilles)
		duree := time.Since(start)
		temps_11_pays2 = append(temps_11_pays2,duree)
		cout2 := glouton.Cout(chemin2)
		chemins =  append(chemins, chemin2)
		exactitude = append(exactitude, math.Abs(cout2-cout)/cout)
	}
	fmt.Println("Comparaison algo approché pour 11 pays", temps_11_pays2, exactitude)

chemin_glouton := chemins[0]
algos_3 := []Algo{    
	{"opt_glouton",opt_glouton.RetireCroisements},

    {"recuit_simule", 	recuit_simule.Recuit},

}
	temps_11_pays3 := make([]time.Duration,0)
exactitude3 := make([]float64,0)

	for _, algo := range algos_3 {
		start := time.Now()
		chemin3 := algo.fonction(chemin_glouton, dicoVilles)
		duree := time.Since(start)
		temps_11_pays3 = append(temps_11_pays3,duree)
		cout3 := glouton.Cout(chemin3)
		exactitude3 = append(exactitude3, math.Abs(cout3-cout)/cout)
	}
	fmt.Println("Comparaison Optimisation pour 11 pays", temps_11_pays3, exactitude3)


}