app.NewWithIDpackage main

import(
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"TSP_general/donnees"
	"TSP_general/dessin_interactif"
)


func main() {

	a := app.New()
	w := a.NewWindow("TSP")
	w.Resize(fyne.NewSize(dessin_interactif.Longueur, dessin_interactif.Largeur))

	ui := dessin_interactif.CreerInterface(donnees.Pays, "FR")

	w.SetContent(ui)
	w.ShowAndRun()
}

