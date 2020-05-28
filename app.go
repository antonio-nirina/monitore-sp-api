package main

import (
	"fmt"
	//"io/ioutil"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
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
	// img, _ := ioutil.ReadFile("output.png")
	res, _ := fyne.LoadResourceFromPath("output.png")
	logo := canvas.NewImageFromResource(res)
	logo.SetMinSize(fyne.NewSize(228, 167))
	p.output = widget.NewLabel("")
	fIt := widget.NewFormItem("text",widget.NewButton("Ok",func(){
		ap.Quit()
	}))
	w := ap.NewWindow("Hello")
	w.Resize(size)
	w.SetMainMenu(createMainMenu())
	w.SetContent(
		widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		p.output,
		logo,
		widget.NewForm(fIt),
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
	helper := fyne.NewMenuItem("Plot",func(){
		p.output.SetText("L'aide ....")
	})
	// Item = append(Item,ItemM,equation,plot)
	menu1 := fyne.NewMenu("Menu",ItemM)
	menu2 := fyne.NewMenu("equation",equation)
	menu3 := fyne.NewMenu("graphique",plot)
	menu4 := fyne.NewMenu("Help",helper)
	main := fyne.NewMainMenu(menu1,menu2,menu3,menu4)

	return main
}
