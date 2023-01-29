package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func ShowPreviewPage() (*fyne.Container, *fyne.Container) {
	img := canvas.NewImageFromFile("desktopClient/2055554.jpg")
	text := canvas.NewText("Word", color.White)

	text.TextStyle = fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: false,
		Symbol:    false,
		TabWidth:  10,
	}
	text.TextSize = 150
	text.Alignment = fyne.TextAlignTrailing

	logo1 := container.New(layout.NewMaxLayout(), img, text)

	text2 := canvas.NewText("Word2", color.White)
	logo2 := container.New(layout.NewMaxLayout(), text2)

	return logo1, logo2
}
