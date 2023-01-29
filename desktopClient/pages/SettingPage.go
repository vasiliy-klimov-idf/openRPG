package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ShowSettings(w fyne.Window) *container.AppTabs {
	var app fyne.App
	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewButton("=light", func() {
			app.Settings().SetTheme(theme.LightTheme())

		})),
		container.NewTabItem("Tab 2", widget.NewButton("=dark", func() {
			app.Settings().SetTheme(theme.DefaultTheme())

		})),
	)

	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewButton("Back", func() {
		menu := ShowMainMenu(w)
		w.SetContent(menu)
	})))
	return tabs
}
