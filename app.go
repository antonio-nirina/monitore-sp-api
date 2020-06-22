package main

import (
	//"io/ioutil"

	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	ly "fyne.io/fyne/layout"
	"fyne.io/fyne/theme"

	// "fyne.io/fyne/layout"
	app1 "fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type Plot struct {
	output *widget.Label
	main   fyne.Widget
}

var p = Plot{}

// var Item []*fyne.MenuItem
var ap = app1.New()
var w = ap.NewWindow("Hello")
var size = fyne.NewSize(500, 200)
var sizeEq = fyne.NewSize(10, 10)

func mainlll() {
	// img, _ := ioutil.ReadFile("output.png")
	p.output = widget.NewLabel("")
	// cnt.Resize(sizeEq)
	main := mainScreen()
	// main.Resize(sizeEq)
	w.Resize(size)
	w.SetMainMenu(createMainMenu())
	w.SetContent(
		main,
	)
	// w.SetFullScreen(true)
	w.ShowAndRun()
}

func createMainMenu() *fyne.MainMenu {
	ItemM := fyne.NewMenuItem("SaveAs", func() {
		widget.NewLabel("Ok ok")
	})
	cut := fyne.NewMenuItem("Cut", func() {
		widget.NewLabel("Ok ok")
	})
	copy := fyne.NewMenuItem("Copy", func() {
		p.output.SetText("Ok ok")
	})
	helper := fyne.NewMenuItem("Preference", func() {
		p.output.SetText("Select theme")
		selected := widget.NewSelect([]string{"LigthTheme", "DarkTheme"}, func(s string) {
			if s == "LigthTheme" {
				ap.Settings().SetTheme(theme.LightTheme())
			} else {
				ap.Settings().SetTheme(theme.DarkTheme())
			}
		})
		back := widget.NewButton("back", func() {
			p.output.SetText("")
			w.SetContent(
				mainScreen())
		})
		back.Style = widget.PrimaryButton
		w.SetContent(
			widget.NewVBox(
				widget.NewLabel("preference"),
				selected,
				widget.NewHBox(back),
			),
		)
	})
	// Item = append(Item,ItemM,equation,plot)
	menu1 := fyne.NewMenu("File", ItemM)
	menu2 := fyne.NewMenu("Edit", copy, cut)
	menu4 := fyne.NewMenu("Help", helper)
	main := fyne.NewMainMenu(menu1, menu2, menu4)

	return main
}

func mainScreen() fyne.Widget {
	pos2 := fyne.NewPos(30, 30)
	pos1 := fyne.NewPos(3, 3)
	res, _ := fyne.LoadResourceFromPath("output.png")
	logo := canvas.NewImageFromResource(res)
	logo.SetMinSize(fyne.NewSize(10, 10))
	logo.Hide()
	abs := widget.NewEntry()
	abs.SetPlaceHolder("Abcisse")
	abs.Move(pos2)
	eq := widget.NewEntry()
	eq.SetPlaceHolder("Equation")
	eq.Move(pos1)
	f := widget.NewVBox(abs, eq, ly.NewSpacer())
	f.Resize(sizeEq)
	cnt := widget.NewTabContainer(widget.NewTabItem("Equation", f))
	btn := widget.NewButton("show", func() {
		fmt.Println("Abcisse:", abs.Text)
		fmt.Println("equation:", eq.Text)
	})
	btn.Style = widget.PrimaryButton
	btnCancel := widget.NewButton("cancel", func() {
		abs.SetText("")
		eq.SetText("")
	})
	btnCancel.Style = widget.DefaultButton
	bx := widget.NewHBox(btn, btnCancel, ly.NewSpacer(), ly.NewSpacer())
	main := widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		p.output,
		logo,
		cnt,
		bx,
	)
	return main
}
