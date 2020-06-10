package main

import (
	// "fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func mainkk() {
	box := tview.NewBox().SetBorder(true).SetTitle("Monitore")
	// box.SetRect(12,12,12,12)
	grid := tview.NewGrid()
	grid.SetTitleColor("Text" tcell.Color17)
	grid.AddItem(box, 2, 5, 3, 2, 0, 0, true)
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
