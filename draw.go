package main

import (
	
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strings"
)

var (
	app *tview.Application
	task *tview.TextView
	taskDetailPane       *tview.Flex
	taskDateDisplay *tview.TextView
	ErrorShow *tview.Flex
	// dateTitle *tview.Flex
	listError *tview.List
	list *tview.List
	test *tview.TextView
	blankCell = tview.NewTextView()
	// statusBar *tview.Pages
)

const pageCount = 2

func trivocli() {
	app = tview.NewApplication()
	task = tview.NewTextView()
	test = tview.NewTextView()
	list = tview.NewList()
	task.SetTitle("[green::b] Data")
	task.SetBorder(true)
	task.SetText("")
	list.SetBorder(true)
	list.AddItem("Text11","show11",0,nil)
	showError()
	test.SetText("ddd")
	// dateTitle.AddItem(test,2,1,false)
	flex := tview.NewFlex()
	flex.AddItem(ErrorShow,30,1,true)
	flex.AddItem(test,10,1,false)
	flex.AddItem(task,0,1,true)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func showError ()  {
	listError = tview.NewList()
	listError.AddItem("","[::d]Id",0,nil)
	listError.AddItem("","Date time",0,nil)
	listError.AddItem("[::d]"+strings.Repeat(string(tcell.RuneS3), 30), "", 0, nil)
	listError.AddItem("Show2","",0,nil)
	ErrorShow = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(listError, 0, 1, true)
	ErrorShow.SetBorder(true)
	ErrorShow.SetTitle("[red::b][::u]E[::-]rrors")
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

