package exp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func gripW() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(500, 500)),
		text1, text2, text3)
	myWindow.SetContent(grid)

	// myWindow.Resize(fyne.NewSize(180, 75))
	myWindow.ShowAndRun()
}
