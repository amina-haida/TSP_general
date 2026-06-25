package recuit_simule

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"TSP_general/opt_glouton"
	"math"
	"math/rand"
	"time"
) 

var lambda = 0.995 

func Opt_aléatoire(trajet []string) []string { 
	rand.Seed(time.Now().UnixNano()) 
	
	n := len(trajet) 
	
	a := rand.Intn(n-1) 
	b := rand.Intn(n-1) 
	
	if a > b { 
		a, b = b, a 
	} 
	newTrajet := make([]string, len(trajet)) 
	copy(newTrajet, trajet) 
	
	for a < b { 
		newTrajet[a], newTrajet[b] = newTrajet[b], newTrajet[a] 
		a++ 
		b--
	} 
	
	return newTrajet 
} 
	
func Recuit(depart string, dicoVilles map[string]donnees.Coordonnees) []string { 
	n := 1000
	trajet := opt_glouton.Fonc_optGlouton(depart, dicoVilles)

	actuel := make([]string, len(trajet)) 
	copy(actuel, trajet) 
	
	meilleur := make([]string, len(trajet)) 
	copy(meilleur, trajet) 
	
	T := 1000.0 
	
	for i := 0; i < n+1; i++ { 
		newTrajet := Opt_aléatoire(actuel) 
	
		coutActuel := glouton.Cout(actuel, dicoVilles) 
		coutNew := glouton.Cout(newTrajet, dicoVilles) 
		
		delta := float64(coutNew - coutActuel) 
		
		if delta < 0 || rand.Float64() < math.Exp(-delta/T) { 
			actuel = newTrajet 
		} 
		
		if glouton.Cout(actuel, dicoVilles) < glouton.Cout(meilleur, dicoVilles) {
			meilleur = actuel 
		}
		
		T *= lambda 
		} 
		
		return meilleur 
}

