package dessin

import (
	"fmt"
	"os"
	"TSP_general/donnees"
)

const marge float64 = 30

var largeur float64 = 1000.0 - 2*marge
var hauteur float64 = 800.0 - 2*marge


func tracer_trajet(
	fichier *os.File,
	objet map[string]donnees.Coordonnees,
	trajet []string,
) {
	for i := 0; i < len(trajet)-1; i++ {

		p1 := objet[trajet[i]]
		p2 := objet[trajet[i+1]]

		min_latitude_objet, max_latitude_objet, min_longitude_objet, max_longitude_objet := donnees.Minmax(objet)

		x1 := ((p1.Longitude-min_longitude_objet)/
			(max_longitude_objet-min_longitude_objet))*
			largeur + marge

		y1 := ((max_latitude_objet-p1.Latitude)/
			(max_latitude_objet-min_latitude_objet))*
			hauteur + marge

		x2 := ((p2.Longitude-min_longitude_objet)/
			(max_longitude_objet-min_longitude_objet))*
			largeur + marge

		y2 := ((max_latitude_objet-p2.Latitude)/
			(max_latitude_objet-min_latitude_objet))*
			hauteur + marge

		fmt.Fprintf(
			fichier,
			`<line x1="%.2f" y1="%.2f" x2="%.2f" y2="%.2f" stroke="red" stroke-width="2"/>`,
			x1, y1, x2, y2,
		)
	}
}



func Dessiner_points(objet map[string]donnees.Coordonnees) {
	fichier, err := os.Create("mon_objet.svg")
	if err != nil {
		panic(err)
	}
	defer fichier.Close()
	fmt.Fprintln(
		fichier,
		`<svg xmlns="http://www.w3.org/2000/svg" width="1000" height="800">`,
	)

	fmt.Fprintln(
		fichier,
		`<rect width="100%" height="100%" fill="white"/>`,
	)

	min_latitude_objet, max_latitude_objet, min_longitude_objet, max_longitude_objet := donnees.Minmax(objet)

	for nom , pays := range objet {
		x := ((pays.Longitude-min_longitude_objet)/(max_longitude_objet-min_longitude_objet))*largeur + marge
		y := ((max_latitude_objet-pays.Latitude)/(max_latitude_objet-min_latitude_objet))*hauteur + marge

		fmt.Fprintf(
			fichier,
			`<circle cx="%.2f" cy="%.2f" r="3" fill="black "/>`,
			x,
			y,
		)

		fmt.Fprintf(
		fichier,
		`<text x="%.2f" y="%.2f" fill="black" font-size="12">%s</text>`+"\n",
		x+5,
		y-5,
		nom,
)

	}

	fmt.Fprintln(fichier, `</svg>`)

}


func Dessiner_trajet(objet map[string]donnees.Coordonnees, trajet []string) {
	fichier, err := os.Create("mon_trajet.svg")
	if err != nil {
		panic(err)
	}
	defer fichier.Close()
	fmt.Fprintln(
		fichier,
		`<svg xmlns="http://www.w3.org/2000/svg" width="1000" height="800">`,
	)

	fmt.Fprintln(
		fichier,
		`<rect width="100%" height="100%" fill="white"/>`,
	)

	min_latitude_objet, max_latitude_objet, min_longitude_objet, max_longitude_objet := donnees.Minmax(objet)

	for nom, pays := range objet {
		x := ((pays.Longitude-min_longitude_objet)/(max_longitude_objet-min_longitude_objet))*largeur + marge
		y := ((max_latitude_objet-pays.Latitude)/(max_latitude_objet-min_latitude_objet))*hauteur + marge

		fmt.Fprintf(
			fichier,
			`<circle cx="%.2f" cy="%.2f" r="3" fill="black "/>`,
			x,
			y,
		)
		fmt.Fprintf(
		fichier,
		`<text x="%.2f" y="%.2f" fill="black" font-size="12">%s</text>`+"\n",
		x+5,
		y-5,
		nom,
		)
	}

	tracer_trajet(fichier, objet, trajet)
	fmt.Fprintln(fichier, `</svg>`)

}
