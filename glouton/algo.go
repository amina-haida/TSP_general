package glouton

import(
	"math"
	"TSP_general/donnees")


func Distance(p1, p2 donnees.Coordonnees) float64 {

	dx := p2.Latitude - p1.Latitude
	dy := p2.Longitude - p1.Longitude

	return dx*dx + dy*dy
}

func Cout(trajet []string) float64 {
	total := 0.0

	for i := 0; i < len(trajet)-1; i++ {
		p1 := donnees.Pays[trajet[i]]
		p2 := donnees.Pays[trajet[i+1]]

		total += math.Sqrt(Distance(p1, p2))
	}

	return total
}

func PlusProche(actuel string,
    nonVisites map[string]donnees.Coordonnees,
    listePays map[string]donnees.Coordonnees) string {

	meilleur := ""
	meilleureDistance := math.Inf(1)

	for nom := range nonVisites {
		d := Distance(listePays[actuel], listePays[nom])

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

		suivant := PlusProche(actuel, nonVisites, pays)

		visites = append(visites, suivant)

		delete(nonVisites, suivant)

		actuel = suivant
	}

	visites = append(visites, depart)

	return visites
}

