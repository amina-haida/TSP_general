package dessin_interactif

import (
    "fyne.io/fyne/v2/app"
)

func Run() {
    a := app.New()
    w := a.NewWindow("TSP")

    w.ShowAndRun()
}