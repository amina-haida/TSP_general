package glouton

import(
	"math"
	"TSP_general/donnees")


func Distance(d1 string, d2 string, listePays map[string]donnees.Coordonnees) float64 {
	p1 := listePays[d1]
	p2 := listePays[d2]

	dx := p2.Latitude - p1.Latitude
	dy := p2.Longitude - p1.Longitude

	return dx*dx + dy*dy
}

func PlusProche(actuel string, nonVisites map[string]donnees.Coordonnees) string {
	meilleur := ""
	meilleureDistance := math.Inf(1)

	for nom := range nonVisites {
		d := Distance(actuel, nom, nonVisites)

		if d < meilleureDistance {
			meilleureDistance = d
			meilleur = nom
		}
	}

	return meilleur
}

func Fonc_glouton(depart string, pays map[string]donnees.Coordonnees) []string {

	visites := []string{depart}

	nonVisites := make(map[string]donnees.Coordonnees)

	for nom, p := range pays {
		if nom != depart {
			nonVisites[nom] = p
		}
	}

	actuel := depart

	for len(nonVisites) > 0 {

		suivant := PlusProche(actuel, nonVisites)

		visites = append(visites, suivant)

		delete(nonVisites, suivant)

		actuel = suivant
	}

	visites = append(visites, depart)

	return visites
}

func Cout(trajet []string, listePays map[string]donnees.Coordonnees) float64 {
	total := 0.0

	for i := 0; i < len(trajet)-1; i++ {
		total += Distance(trajet[i], trajet[i+1], listePays)
	}

	return total
}