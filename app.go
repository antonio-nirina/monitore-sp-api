package main

import (
	"fmt"

	"fyne.io/fyne"
	// "fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)

type Plot struct {
	output  *widget.Label
	window  fyne.Window
}

var p = Plot{}
// var Item []*fyne.MenuItem
var ap = app.New()

func main() {
	size := fyne.NewSize(500,200)
	// sizeIn := fyne.NewSize(50,20)
	p.output = widget.NewLabel("")
	w := ap.NewWindow("Hello")
	w.Resize(size)
	w.SetMainMenu(createMainMenu())
	w.SetContent(
		widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		p.output,
		/*widget.NewButton("Quit", func() {
			ap.Quit()
		}),*/
	))

	w.ShowAndRun()
}

func createMainMenu() *fyne.MainMenu {
	ItemM := fyne.NewMenuItem("Cours",func(){
		widget.NewLabel("Ok ok")
	})
	equation := fyne.NewMenuItem("Equation",func(){
		p.output.SetText("Ok ok")
	})
	plot := fyne.NewMenuItem("Plot",func(){
		fmt.Println("Ok Ok")
	})
	// Item = append(Item,ItemM,equation,plot)
	menu1 := fyne.NewMenu("Show",ItemM)
	menu2 := fyne.NewMenu("equation",equation)
	menu3 := fyne.NewMenu("graphique",plot)
	main := fyne.NewMainMenu(menu1,menu2,menu3)

	return main
}
