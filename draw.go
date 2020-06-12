package main

import (
	"github.com/rivo/tview"
)

var (
	app *tview.Application
	// statusBar *tview.Pages
)

const pageCount = 2

func main() {
	app = tview.NewApplication()
	flex := tview.NewFlex().AddItem(
		tview.NewBox().SetBorder(true).SetTitle("[red::b] List Error"), 40, 1, false)
	flex.AddItem(tview.NewBox().SetBorder(true).SetTitle("[green::b] Data"), 0, 1, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
