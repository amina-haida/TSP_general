package dessininteractif

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"TSP_general/donnees"
	)

func Dessin_points(Points map[string]donnees.Coordonnees) {
	a := app.New()
	w := a.NewWindow("TSP")

	w.Resize(fyne.NewSize(1000, 600))
	minLat, maxLat, minLon, maxLon := donnees.Minmax(Points)


	content := container.NewWithoutLayout()
	for ville, coord := range Points { 
		circle := canvas.NewCircle(color.Black)

		x := ((coord.Longitude - minLon) / (maxLon - minLon)) * 1000.0
		y := ((coord.Latitude - minLat) / (maxLat - minLat)) * 600
		
		circle.Move(fyne.NewPos(float32(x-3), float32(y-3)))
		circle.Resize(fyne.NewSize(6, 6))

		label := canvas.NewText(ville, color.Black)
		label.Move(fyne.NewPos(float32(x+8), float32(y)))

		content.Add(circle)
		content.Add(label)
		content.Refresh()
	}	
	w.SetContent(content)
	w.ShowAndRun()
}