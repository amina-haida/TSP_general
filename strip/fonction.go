package strip

import (
	"TSP_general/donnees"
	"TSP_general/glouton"
	"fmt"
)


func Combiner(L1,L2 []string, Points map[string]donnees.Coordonnees)[]string{

	i1:=0
	i2:=0
	var L []string

	for i:=0 ; i<len(L1)+len(L2) ; i++{

		if i1 == len(L1) {
			L=append(L, L2[i2])
			i2++

		} else if i2==len(L2){
			L=append(L, L1[i1])

			i1++
		}else if Points[L1[i1]].Latitude < Points[L2[i2]].Latitude{

			L=append(L,L1[i1] )

			i1++
		} else {

			L = append(L, L2[i2])
			i2++
		}
	}
	return L
}

func Tri_latitude(listeVille []string, Points map[string]donnees.Coordonnees) []string{

	if len(listeVille) <2 {
		return listeVille
	}else {
		milieu := len(listeVille)/2
		Liste1 := listeVille[:milieu]
		Liste2 := listeVille[milieu:]

		L1_triee := Tri_latitude(Liste1, Points)
		L2_triee := Tri_latitude(Liste2, Points)

		L_finale := Combiner(L1_triee, L2_triee, Points)

		return L_finale

	}
}

func Inverser(l []string) {
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
}

func Strip(Points map[string]donnees.Coordonnees, depart string, mesure float64) []string {

	_, _, minLon, maxLon := donnees.Minmax(Points)

	var trajet []string

	nbBandes := len(Points) / 2

	largeur := (maxLon - minLon) / float64(nbBandes)

	sens := true

	for x := minLon - mesure; x <= maxLon; x += largeur {

		var bande []string

		for ville := range Points {

			lon := Points[ville].Longitude

			if lon >= x && lon < x+largeur {
				bande = append(bande, ville)
			}
		}

		bande = Tri_latitude(bande, Points)

		if !sens {
			Inverser(bande)
		}

		trajet = append(trajet, bande...)

		sens = !sens
	}

	pos := -1

	for i, v := range trajet {
		if v == depart {
			pos = i
			break
		}
	}

	if pos != -1 {
		trajet = append(trajet[pos:], trajet[:pos]...)
	}

	if len(trajet) > 0 {
		trajet = append(trajet, trajet[0])
	}

	return trajet
}

func MeilleurStrip(Points map[string]donnees.Coordonnees, depart string) []string {
	min := Strip(Points, depart, 0)
	coutMin := glouton.Cout(min)

	_, _, minLon, maxLon := donnees.Minmax(Points)

	nbBandes := len(Points) / 2
	largeur := (maxLon - minLon) / float64(nbBandes) 

	for k := 0; k < nbBandes; k++ {

		mesure := largeur * float64(k) / float64(nbBandes)


		trajet := Strip(Points, depart, mesure)

		cout := glouton.Cout(trajet)

		fmt.Println(trajet)

		if cout < coutMin {
			min = trajet
			coutMin = cout
		}
	}
	return min
}