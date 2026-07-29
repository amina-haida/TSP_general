package donnees

import (
	"testing"
)




func Test_Ville_aleatoire(t *testing.T){

	map_pays := Ville_aleatoire(27)
	longueur:= 27
if longueur != len(map_pays){
	t.Fatalf(
            "nombre de villes incorrect : obtenu %d, attendu %d",
            len(map_pays),
            27,
        )}}