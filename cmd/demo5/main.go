package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("Hello world data binding")
	w := a.NewWindow("Bind Name")

	name := binding.BindPreferenceString("user.name", a.Preferences())
	w.SetContent(container.NewHBox(
		widget.NewLabel("Please enter your name:"),
		widget.NewEntryWithData(name)),
	)

	w.ShowAndRun()
}
