package main

import (
	//"io/ioutil"

	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"

	// "fyne.io/fyne/layout"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type Plot struct {
	output *widget.Label
	window fyne.Window
}

var p = Plot{}

// var Item []*fyne.MenuItem
var ap = app.New()

func mainff() {
	size := fyne.NewSize(500, 200)
	sizeEq := fyne.NewSize(10, 10)
	pos2 := fyne.NewPos(30, 30)
	pos1 := fyne.NewPos(3, 3)
	// img, _ := ioutil.ReadFile("output.png")
	res, _ := fyne.LoadResourceFromPath("output.png")
	logo := canvas.NewImageFromResource(res)
	logo.SetMinSize(fyne.NewSize(10, 10))
	logo.Hide()
	p.output = widget.NewLabel("")
	abs := widget.NewEntry()
	abs.SetPlaceHolder("Abcisse")
	abs.Move(pos2)
	eq := widget.NewEntry()
	eq.SetPlaceHolder("Equation")
	eq.Move(pos1)
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
	f := widget.NewVBox(abs, eq, layout.NewSpacer())
	f.Resize(sizeEq)
	bx := widget.NewHBox(btn, btnCancel, layout.NewSpacer(), layout.NewSpacer())
	cnt := widget.NewTabContainer(widget.NewTabItem("Equation", f))
	cnt.Resize(sizeEq)
	fmt.Println(cnt.Size())
	w := ap.NewWindow("Hello")
	w.Resize(size)
	w.SetMainMenu(createMainMenu())
	w.SetContent(
		widget.NewVBox(
			widget.NewLabel("Hello Fyne!"),
			p.output,
			logo,
			cnt,
			bx,
		),
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
		p.output.SetText("L'aide ....")
	})
	// Item = append(Item,ItemM,equation,plot)
	menu1 := fyne.NewMenu("File", ItemM)
	menu2 := fyne.NewMenu("Edit", copy, cut)
	menu4 := fyne.NewMenu("Help", helper)
	main := fyne.NewMainMenu(menu1, menu2, menu4)

	return main
}
