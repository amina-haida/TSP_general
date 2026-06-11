package opt_glouton

import (
	"TSP_general/donnees"

)


func orientation(a, b, c donnees.Coordonnees) float64 {
	return (b.Latitude-a.Latitude)*(c.Longitude-a.Longitude) -
		(b.Longitude-a.Longitude)*(c.Latitude-a.Latitude)
}

func croisement(a, b, c, d donnees.Coordonnees) bool {
	o1 := orientation(a, b, c)
	o2 := orientation(a, b, d)
	o3 := orientation(c, d, a)
	o4 := orientation(c, d, b)

	return (o1*o2 < 0) && (o3*o4 < 0)
}

func echange(tour []string, i, j int) {
	for g, d := i+1, j; g < d; g, d = g+1, d-1 {
		tour[g], tour[d] = tour[d], tour[g]
	}
}

func RetireCroisements(trajet []string, pays map[string]donnees.Coordonnees) []string {

	n := len(trajet)

	amelio := true

	for amelio {

		amelio = false

		for i := 0; i < n-3; i++ {

			a := pays[trajet[i]]
			b := pays[trajet[i+1]]

			for j := i + 2; j < n-1; j++ {

				if i == 0 && j == n-2 {
					continue
				}

				c := pays[trajet[j]]
				d := pays[trajet[j+1]]

				if croisement(a, b, c, d) {

					echange(trajet, i, j)

					amelio = true
					break
				}
			}

			if amelio {
				break
			}
		}
	}

	return trajet
}