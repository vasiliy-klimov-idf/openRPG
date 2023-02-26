package page

import (
	"client/core"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
)

func EscapeMenu() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := widgets.NewParagraph()
	header.Text = "Select your Class."
	header.SetRect(0, 0, 50, 0)
	header.Border = false
	header.TextStyle.Fg = ui.ColorGreen
	header.TextStyle.Bg = ui.ColorBlue
	header.TextStyle.Modifier = ui.ModifierBold

	l := widgets.NewList()
	l.Title = "Menu Pause"
	l.Rows = []string{"- [Save Game] \n", "- [Load Game] \n", "- [History] \n", "- [Help] \n", "- [Custom] \n", "- [EXIT] \n"}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 5, 25, 8)
	ui.Render(l)

	p := widgets.NewParagraph()
	p.Text = ""
	p.Border = false

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/4, p),
			ui.NewCol(1.0/4, p),
			ui.NewCol(1.0/4, p),
		),
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/4, p),
			ui.NewCol(1.0/4, l),
			ui.NewCol(1.0/4, p),
		),
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/4, p),
			ui.NewCol(1.0/4, p),
			ui.NewCol(1.0/4, p),
		),
	) //дописать внизу инфо о кнопках

	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Enter>":
				//select to class object
			case "p":
				//Sorry for this code. IDK how to fix this;
				switch core.PageNumber {
				case 2:
					SelectRaceMenu()
				case 3:
					SelectClassMenu()
				}
			case "<Down>":
				l.ScrollDown()
			case "<Up>":
				l.ScrollUp()
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		default:
			ui.Render(grid)
		}
	}
}
