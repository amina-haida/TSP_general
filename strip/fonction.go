package strip

import (
	"TSP_general/donnees"
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

func Strip(Points map[string]donnees.Coordonnees) []string{
	_, _, minLon, maxLon := donnees.Minmax(Points)
	var trajet []string
	n := float64(len(Points)/2)

	for i := minLon; i < maxLon; i = i + n{
		var trajet_tempo []string
		for ville := range Points{
			if i < Points[ville].Longitude && Points[ville].Longitude < i + n{
			trajet_tempo = append(trajet_tempo, ville)
			}
		}
		Tri_latitude(trajet_tempo, Points)

		for _, ville_tri := range trajet_tempo{
			trajet = append(trajet, ville_tri)
		}
	}
	return trajet	
}

