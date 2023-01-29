package main

import (
	"desktop_client/pages"
	"fyne.io/fyne/v2/app"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Max Layout")

	logo1, logo2 := pages.ShowPreviewPage()
	mainMenu := pages.ShowMainMenu(w)

	go func() {
		w.SetContent(logo1)
		logo1.Refresh()
		time.Sleep(1 * time.Second)
		w.SetContent(logo2)
		logo2.Refresh()
		time.Sleep(1 * time.Second)
		w.SetContent(mainMenu)
		w.Show()
		logo1.RemoveAll()
		logo2.RemoveAll()

	}()
	w.Show()
	a.Run()
}
