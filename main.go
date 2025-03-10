package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Hello Fyne")

	a := app.New()
	window := a.NewWindow("Test Fyne")
	label := widget.NewLabel("Hello Fyne!")
	button := widget.NewButton("Hi!", func() {
		fmt.Println("Button Pressed!")
	})

	window.SetContent(
		container.NewVBox(label, button),
	)
	window.ShowAndRun()
}
