package main

import "fyne.io/fyne/v2/app"

func main() {
	a := app.NewWithID("io.fyne.workshop.todo")
	w := a.NewWindow("TODO")

	w.SetContent(makeUI())

	w.ShowAndRun()
}
