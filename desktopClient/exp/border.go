package exp

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func Border() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	bottom := canvas.NewText("bottom", color.White)
	right := canvas.NewText("right", color.White)
	content := container.NewBorder(top, bottom, left, right)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
