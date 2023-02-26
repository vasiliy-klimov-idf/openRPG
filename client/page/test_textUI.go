package page

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
)

func TestTextPage() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	l := widgets.NewList()
	l.Title = "List"
	l.Rows = []string{
		"[0] github.com/gizak/termui/v3",
		"[1] [你好，世界](fg:blue)",
		"[2] [こんにちは世界](fg:red)",
		"[3] [color](fg:white,bg:green) output",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] foo",
		"[8] bar",
		"[9] baz",
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 25, 8)

	l2 := widgets.NewList()
	l2.Title = "List"
	l2.Rows = []string{
		"[0] github.com/gizak/termui/v3",
		"[1] [你好，世界](fg:blue)",
		"[2] [こんにちは世界](fg:red)",
		"[3] [color](fg:white,bg:green) output",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] foo",
		"[8] bar",
		"[9] baz",
	}
	l2.TextStyle = ui.NewStyle(ui.ColorYellow)
	l2.WrapText = false
	l2.SetRect(0, 0, 25, 8)

	p := widgets.NewParagraph()
	p.Title = "Demonstration"

	p2 := widgets.NewParagraph()
	p2.Title = "Demonstration2"

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/1, p),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/4, l),
			ui.NewCol(1.0/4, l2),
			ui.NewCol(1.0/2, p2),
		),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	var list *widgets.List
	var paragraph *widgets.Paragraph
	var num int
	list = l
	ArrayOfList := []*widgets.List{l, l2}
	ArrayOfParag := []*widgets.Paragraph{p, p2}
	paragraph = p
	//list.SelectedRowStyle
	for {
		select {
		case e := <-uiEvents:
			if e.ID == "<Right>" {
				num++
				if num > len(ArrayOfList)-1 {
					num = len(ArrayOfList) - 1
				}
				for i, ol := range ArrayOfList {
					for pi, op := range ArrayOfParag {
						if num == i && num > 0 && num == pi {
							ArrayOfList[i-1].TextStyle = ui.NewStyle(ui.ColorYellow)
							ArrayOfList[i-1].SelectedRowStyle = ui.Style{ui.ColorYellow, ui.ColorClear, ui.ModifierBold}
							ArrayOfList[i].TextStyle = ui.NewStyle(ui.ColorRed)
							ArrayOfList[i].SelectedRowStyle = ui.Style{ui.ColorGreen, ui.ColorRed, ui.ModifierBold}
							list = ol
							paragraph = op
						}
					}
				}
			}
			if e.ID == "<Left>" {
				num--
				if num < 0 {
					num = 0
				}
				for i, ol := range ArrayOfList {
					for pi, op := range ArrayOfParag {
						if num == i && num < len(ArrayOfList)-1 && num == pi && num < len(ArrayOfParag)-1 {
							ArrayOfList[i+1].TextStyle = ui.NewStyle(ui.ColorYellow)
							ArrayOfList[i+1].SelectedRowStyle = ui.Style{ui.ColorYellow, ui.ColorClear, ui.ModifierBold}
							ArrayOfList[i].TextStyle = ui.NewStyle(ui.ColorRed)
							ArrayOfList[i].SelectedRowStyle = ui.Style{ui.ColorGreen, ui.ColorRed, ui.ModifierBold}
							list = ol
							paragraph = op
						}
					}
				}
			}

			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				list.SelectedRowStyle = ui.Style{ui.ColorGreen, ui.ColorRed, ui.ModifierBold}
				paragraph.Text = list.Rows[list.SelectedRow]
				list.ScrollDown()
			case "k", "<Up>":
				list.SelectedRowStyle = ui.Style{ui.ColorGreen, ui.ColorRed, ui.ModifierBold}
				paragraph.Text = list.Rows[list.SelectedRow]
				list.ScrollUp()
			case "<Resize>":
				x, y := ui.TerminalDimensions()
				list.SetRect(0, 0, x, y)
			}
		default:
			ui.Render(grid)
			list.SelectedRowStyle = ui.Style{ui.ColorGreen, ui.ColorRed, ui.ModifierBold}
			list.TextStyle = ui.NewStyle(ui.ColorRed)
			paragraph.Text = list.Rows[list.SelectedRow]
		}
	}
}
