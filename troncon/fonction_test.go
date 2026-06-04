package troncon

import (
	"testing"
)

func Test_troncons_rec2(t *testing.T) {

	Reseau_ferre = []Troncon{
		{"1", "2"},
		{"2", "3"},
		{"2", "4"},
		{"1", "5"},
		{"1", "6"},
	}

	var chemin_test []string

	Troncons_rec2("1", "1", &chemin_test)

	res_attendu := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}
	for index, _ := range chemin_test {
		if chemin_test[index] != res_attendu[index] {
			t.Errorf("Erreur à l'indice%d", index)
		}

	}
}

func Test_chercher_ville(t *testing.T) {
	Reseau_ferre = []Troncon{
		{"1", "2"},
		{"2", "3"},
		{"2", "4"},
		{"1", "5"},
		{"1", "6"},
	}

	var chemin_test []string

	chemin_attendu := []string{"4", "2", "1"}

	Chercher_ville("1", "1", "4", &chemin_test)

	for index := range chemin_test {
		if chemin_test[index] != chemin_attendu[index] {
			t.Errorf("Erreur à l'indice %d", index)
		}
	}

}
