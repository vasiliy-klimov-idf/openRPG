package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowMainMenu(w fyne.Window) *fyne.Container {

	startBottom := widget.NewButton("Start Game", func() {

	})
	loadBottom := widget.NewButton("Load", func() {

	})
	settingsBottom := widget.NewButton("Settings", func() {
		settingsCont := ShowSettings(w)
		w.SetContent(settingsCont)
	})
	exitBottom := widget.NewButton("Exit", func() {

	})
	buttonBox := container.NewVBox(startBottom, loadBottom, settingsBottom, exitBottom)

	mainGrid := container.NewBorder(layout.NewSpacer(), nil, buttonBox, nil)
	return mainGrid
}
