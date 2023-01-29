package exp

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("win")

	imageA := canvas.NewImageFromFile("desktopClient/2055554.jpg")
	imageA.FillMode = canvas.ImageFillStretch

	imageB := canvas.NewImageFromFile("desktopClient/b.jpg")
	imageB.FillMode = canvas.ImageFillOriginal

	var contentA *fyne.Container
	var contentB *fyne.Container
	buttonA := widget.NewButton("change to B", func() {
		fmt.Println("xxx")
		window.SetContent(contentB)
		window.Show()
	})
	buttonB := widget.NewButton("change to A", func() {
		fmt.Println("yyy")
		window.SetContent(contentA)
		window.Show()

	})
	contentA = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), imageA, buttonA)
	contentB = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), imageB, buttonB)

	window.SetContent(contentA)
	window.ShowAndRun()
}
