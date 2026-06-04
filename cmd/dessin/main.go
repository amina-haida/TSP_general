package main

import (
	"TSP_general/donnees"
	"TSP_general/dessin"
)

func main() {
	europe := donnees.Europe()
	trajet := []string{
    "FR", "GI", "MT", "PT", "ES", "AD", "VA", "IT", "MC", "GR",
    "SM", "AL", "ME", "MK", "BA", "XK", "CH", "HR", "LI", "CY",
    "SI", "RS", "JE", "GG", "AT", "BG", "LU", "BE", "HU", "CZ",
    "DE", "RO", "NL", "SK", "TR", "IE", "IM", "MD", "PL", "GB",
    "DK", "UA", "LT", "BY", "NO", "LV", "FO", "SE", "EE", "FI",
    "FR",
}

	dessin.Dessiner_points(europe)

	dessin.Dessiner_trajet(europe, trajet)

}