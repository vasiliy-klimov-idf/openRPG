package page

import (
	"client/core"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	r "github.com/vasiliy-klimov-idf/openrpg_core/src/races"
	"log"
)

func SelectRaceMenu() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := widgets.NewParagraph()
	header.Text = "Select your Race."
	header.SetRect(0, 0, 50, 0)
	header.Border = false
	header.TextStyle.Fg = ui.ColorGreen
	header.TextStyle.Bg = ui.ColorBlue
	header.TextStyle.Modifier = ui.ModifierBold

	l := widgets.NewList()
	l.Title = "List"
	l.Rows = r.RaceArrayName
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 5, 25, 8)
	ui.Render(l)

	p := widgets.NewParagraph()
	p.Title = "Demonstration"

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/13, header),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/4, l),
			ui.NewCol(1.0/2, p),
		),
	) //дописать внизу инфо о кнопках
	ui.Render(grid)

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Enter>":
				for _, race := range r.RaceArray {
					if race.Name == l.Rows[l.SelectedRow] {
						core.SelectedRace = race.Name
						return
					}
				}
			case "<Down>":
				l.ScrollDown()
				for _, race := range r.RaceArray {
					if race.Name == l.Rows[l.SelectedRow] {
						//info, _ := json.MarshalIndent(class, "", "  ")
						p.Text = "Name: " + race.Name + "\n" +
							"Description: " + race.Description + "\n"
					}
				}
			case "<Up>":
				l.ScrollUp()
				for _, race := range r.RaceArray {
					if race.Name == l.Rows[l.SelectedRow] {
						//info, _ := json.MarshalIndent(class, "", "  ")
						p.Text = "Name: " + race.Name + "\n" +
							"Description: " + race.Description + "\n"
					}
				}
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 5, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		default:
			ui.Render(grid)
		}
	}
}
