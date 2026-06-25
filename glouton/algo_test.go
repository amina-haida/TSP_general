package glouton


import (
	"testing"
	"TSP_general/donnees"
)

func TestDistance(t *testing.T) {

	d := Distance(donnees.Monde["FR"], donnees.Monde["FO"])

	expected := 328.667885068034

	if d != expected {
		t.Errorf("distance incorrecte, obtenu %v, attendu %v", d, expected)
	}
}

func TestGlouton(t *testing.T) {

	p := map[string]donnees.Coordonnees{
		"FR": {Latitude: 0, Longitude: 0},
		"DE": {Latitude: 1, Longitude: 1},
		"ES": {Latitude: 2, Longitude: 2},
	}

	tour := Fonc_glouton("FR", p)

	if tour[0] != tour[len(tour)-1] {
		t.Error("le cycle ne revient pas au départ")
	}

	if len(tour) != len(p)+1 {
		t.Errorf("taille incorrecte : %d", len(tour))
	}
}

func TestCout(t *testing.T) {

	
	cycle := []string{"FR", "GA", "RU", "ZM"}

	c := Cout(cycle, donnees.Monde)

	if c <= 0 {
		t.Error("le coût doit être positif")
	}
}

func TestEurope(t *testing.T) {

	europe := donnees.Europe()

	if len(europe) == 0 {
		t.Error("Europe vide")
	}

	for _, p := range europe {
		if p.Longitude < -10 || p.Longitude > 40 || p.Latitude < 35 || p.Latitude > 70 {
			t.Error("pays hors Europe détecté")
		}
	}
}
