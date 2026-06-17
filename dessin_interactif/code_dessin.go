package dessin_interactif

import (
	"TSP_general/donnees"
	"TSP_general/dynamique"
	"TSP_general/glouton"
	"TSP_general/opt_glouton"
	"TSP_general/recuit_simule"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

func CreerInterface(Points map[string]donnees.Coordonnees, départ string) fyne.CanvasObject {

	afficheTrajet := container.NewWithoutLayout()
	affichePoints := container.NewWithoutLayout()

	Dessin_points(Points, affichePoints)

	circle := canvas.NewCircle(color.Black)

		x, y := coordonnees(Points[départ], Points)

		circle.Move(fyne.NewPos(float32(x-5), float32(y-5)))
		circle.Resize(fyne.NewSize(10, 10))

		label := canvas.NewText(départ, color.Black)
		label.Move(fyne.NewPos(float32(x+8), float32(y)))

		affichePoints.Add(circle)

	var trajet []string


	zoomPlus := widget.NewButton("+", func() {

	Zoom *= 1.2

	Redessiner(
		Points,
		départ,
		trajet,
		affichePoints,
		afficheTrajet,
	)
	})

	zoomMoins := widget.NewButton("-", func() {

	Zoom /= 1.2

	Redessiner(Points, départ, trajet, affichePoints, afficheTrajet)
	})

	
	gauche := widget.NewButton("←", func() {
		OffsetX += 50
		Redessiner(Points, départ, trajet, affichePoints, afficheTrajet)
	})

	droite := widget.NewButton("→", func() {
		OffsetX -= 50
		Redessiner(Points, départ, trajet, affichePoints, afficheTrajet)
	})

	haut := widget.NewButton("↑", func() {
		OffsetY += 50
		Redessiner(Points, départ, trajet, affichePoints, afficheTrajet)
	})

	bas := widget.NewButton("↓", func() {
		OffsetY -= 50
		Redessiner(Points, départ, trajet, affichePoints, afficheTrajet)
	})

	selection := widget.NewSelect(
		[]string{"Glouton", "2-opt", "Recuit simulé", "Dynamique"},
		func(choix string) {
			afficheTrajet.RemoveAll()

			if choix == "Glouton"{
				trajet = glouton.Fonc_glouton(départ, Points)
			}
			if choix == "2-opt"{
				trajet_ini := glouton.Fonc_glouton(départ, Points)
				trajet = opt_glouton.RetireCroisements(trajet_ini, Points)
			}
			if choix == "Recuit simulé"{
				trajet_ini := glouton.Fonc_glouton(départ, Points)
				trajetOpt := opt_glouton.RetireCroisements(trajet_ini, Points)
				trajet = recuit_simule.Recuit(trajetOpt, 1000)
			}
			if choix == "Dynamique"{
				trajet = dynamique.TSP_dynamique(Points, départ)
						}
			Redessiner(Points, départ, trajet, affichePoints, afficheTrajet)
		},
	)
	topBar := container.NewHBox(
		selection,
		zoomPlus,
		zoomMoins,
		gauche,
		droite,
		haut,
		bas,
	)

	dessin := container.NewWithoutLayout(afficheTrajet, affichePoints)

	return container.NewBorder(
		topBar,
		nil,
		nil,
		nil,
		dessin,
	)
}
