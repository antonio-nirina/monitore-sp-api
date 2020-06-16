package main

import (
	
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	app *tview.Application
	task *tview.TextView
	taskDetailPane       *tview.Flex
	taskDateDisplay *tview.TextView
	blankCell = tview.NewTextView()
	// statusBar *tview.Pages
)

const pageCount = 2

func main() {
	app = tview.NewApplication()
	task = tview.NewTextView().SetDynamicColors(true)
	taskDateDisplay = tview.NewTextView()
	flex := tview.NewFlex()
	flex.AddItem(taskDateDisplay.SetText("Error"),0,1,false).
		SetBorder(true).SetTitle("Error")
		//tview.NewBox().SetBorder(true).SetTitle("[red::b] List Error"), 40, 1, false)
	// flex.AddItem(tview.NewBox().SetBorder(true).SetTitle("[green::b] Data"), 0, 1, false)
	flex.AddItem(task.SetText("Text"),0,1,false).SetBorder(true).SetTitle("Data")
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func makeDateRow() *tview.Flex {
	task = tview.NewTextView().SetDynamicColors(true)
	taskDetailPane = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(task, 2, 1, true).
		AddItem(blankCell, 1, 1, false).
		AddItem(makeDateRow(), 1, 1, true).
		AddItem(blankCell, 1, 1, false)
	taskDetailPane.SetBorder(true).SetTitle("Task Detail")
	taskDetailPane.AddItem(tview.NewBox().SetBorder(true).SetTitle("[green::b] Data"), 0, 1, false)
	taskDateDisplay = tview.NewTextView().SetDynamicColors(true)
	taskDate := tview.NewTextView().SetText("Task Not[::u]e[::-]:").SetTextColor(tcell.ColorDimGray)
	return tview.NewFlex().
		AddItem(taskDateDisplay, 0, 2, true).
		AddItem(taskDate, 14, 0, true)
}

