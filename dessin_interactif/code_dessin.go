package dessin_interactif

import (
	"TSP_general/donnees"
	"TSP_general/dynamique"
	"TSP_general/glouton"
	"TSP_general/opt_glouton"
	"TSP_general/recuit_simule"
	"TSP_general/strip"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const Longueur = 1000.0
const Largeur = 600.0

var Zoom float64 = 1.0

var OffsetX float64 = 0
var OffsetY float64 = 0

const CentreX = Longueur / 2
const CentreY = Largeur / 2

const Marge = 40.0

func coordonnees(a donnees.Coordonnees, Points map[string]donnees.Coordonnees) (x, y float64) {
	minLat, maxLat, minLon, maxLon := donnees.Minmax(Points)

	newLongueur := (Longueur - 2*Marge) 
	newLargeur := (Largeur - 2*Marge) 

	x0 := Marge +
    ((a.Longitude-minLon)/(maxLon-minLon))*newLongueur

	y0 := Largeur - Marge -
    ((a.Latitude-minLat)/(maxLat-minLat))*newLargeur

	x = CentreX + (x0-CentreX)*Zoom + OffsetX
	y = CentreY + (y0-CentreY)*Zoom + OffsetY

	return
}




func Dessin_points(Points map[string]donnees.Coordonnees, content *fyne.Container) {

	for ville, coord := range Points { 
		circle := canvas.NewCircle(color.RGBA{
			R: 220,
			G: 50,
			B: 50,
			A: 255,
		})

		x, y := coordonnees(coord, Points)

		circle.Move(fyne.NewPos(float32(x-3), float32(y-3)))
		circle.Resize(fyne.NewSize(6, 6))

		label := canvas.NewText(ville, color.Black)
		label.Move(fyne.NewPos(float32(x+8), float32(y)))

		content.Add(circle)
		content.Add(label)
	}	
}

func Dessin_trajet(Points map[string]donnees.Coordonnees, trajet []string, content *fyne.Container) {

	for i := 0; i < len(trajet)-1; i++ {
		a := Points[trajet[i]]
		b := Points[trajet[i+1]]

		x_a, y_a := coordonnees(a, Points)
		x_b, y_b := coordonnees(b, Points)

		line := canvas.NewLine(color.RGBA{
			R: 0,
			G: 90,
			B: 220,
			A: 255,
		})

		line.StrokeWidth = 2

		line.Position1 = fyne.NewPos(float32(x_a), float32(y_a))
		line.Position2 = fyne.NewPos(float32(x_b), float32(y_b))

		line.StrokeWidth = 2

		content.Add(line)
		
	}
}

func Redessiner(Points map[string]donnees.Coordonnees, depart string, trajet []string, affichePoints *fyne.Container, afficheTrajet *fyne.Container) {

	affichePoints.RemoveAll()
	afficheTrajet.RemoveAll()

	Dessin_points(Points, affichePoints)

	x, y := coordonnees(Points[depart], Points)

	circle := canvas.NewCircle(color.Black)
	circle.Move(fyne.NewPos(float32(x-5), float32(y-5)))
	circle.Resize(fyne.NewSize(10, 10))

	affichePoints.Add(circle)

	Dessin_trajet(Points, trajet, afficheTrajet)
	

	affichePoints.Refresh()
	afficheTrajet.Refresh()
}

func CreerInterface(w fyne.Window,) fyne.CanvasObject {

	afficheTrajet := container.NewWithoutLayout()
	affichePoints := container.NewWithoutLayout()

	var Points map[string]donnees.Coordonnees
	var selectionDepart *widget.Select
	var depart string
	var trajet []string

	updateVilles := func() {
		villes := make([]string, 0, len(Points))
		for v := range Points {
			villes = append(villes, v)
		}

		selectionDepart.Options = villes
		selectionDepart.SetSelected("")
		selectionDepart.Refresh()
	}
	distanceLabel := widget.NewLabel("Distance : ---")
	selectionPays := widget.NewSelect(
		[]string{
			"Europe",
			"Monde",
			"Aléatoire",
		},
		func(choix string) {
			affichePoints.RemoveAll()

			if choix == "Europe"{
				Points = donnees.Europe()
			}
			if choix == "Monde"{
				Points = donnees.Monde
			}
			if choix == "Aléatoire"{
				Points = donnees.GenererVilles(15)
			}
			Dessin_points(Points, affichePoints)

			updateVilles()
		},
	)


	selectionDepart = widget.NewSelect([]string{}, func(choix string) {
		depart = choix
	})

	selectionAlgo := widget.NewSelect(
		[]string{"Glouton", "2-opt", "Recuit simulé", "Dynamique", "Strip"},
		func(choix string) {
			afficheTrajet.RemoveAll()

			if choix == "Glouton"{
				trajet = glouton.Fonc_glouton(depart, Points)
			}
			if choix == "2-opt"{
				trajet = opt_glouton.Fonc_optGlouton(depart, Points)
			}
			if choix == "Recuit simulé"{
				trajet = recuit_simule.Recuit(depart, Points)
			}
			if choix == "Dynamique"{
				if len(Points) < 25{
					trajet = dynamique.TSP_dynamique(depart, Points)
				}else{
					dialog.ShowInformation(
						"Erreur",
						"Programmation dynamique limitée à 25 villes",
						w,
    				)
				}
			}
			if choix == "Strip"{
				trajet = strip.MeilleurStrip(depart, Points)
			}
			distanceLabel.SetText(
			fmt.Sprintf(
				"Distance : %.0f km",
				glouton.Cout(trajet, Points),
				),
			)
			Redessiner(Points, depart, trajet, affichePoints, afficheTrajet)

		},
	)


	zoomPlus := widget.NewButton("+", func() {

	Zoom *= 1.2

	Redessiner(
		Points,
		depart,
		trajet,
		affichePoints,
		afficheTrajet,
	)
	})

	zoomMoins := widget.NewButton("-", func() {

	Zoom /= 1.2

	Redessiner(Points, depart, trajet, affichePoints, afficheTrajet)
	})

	
	gauche := widget.NewButton("←", func() {
		OffsetX += 50
		Redessiner(Points, depart, trajet, affichePoints, afficheTrajet)
	})

	droite := widget.NewButton("→", func() {
		OffsetX -= 50
		Redessiner(Points, depart, trajet, affichePoints, afficheTrajet)
	})

	haut := widget.NewButton("↑", func() {
		OffsetY += 50
		Redessiner(Points, depart, trajet, affichePoints, afficheTrajet)
	})

	bas := widget.NewButton("↓", func() {
		OffsetY -= 50
		Redessiner(Points, depart, trajet, affichePoints, afficheTrajet)
	})
	topBar := container.NewHBox(
		selectionPays,
		selectionDepart,
		selectionAlgo,
		zoomPlus,
		zoomMoins,
		gauche,
		droite,
		haut,
		bas,
	)


	background := canvas.NewRectangle(color.White)
	background.Resize(fyne.NewSize(1000, 600))

	dessin := container.NewWithoutLayout(
		background,
		afficheTrajet,
		affichePoints,
	)
	

	return container.NewBorder(
		topBar,
		distanceLabel,
		nil,
		nil,
		dessin,
	)
}
