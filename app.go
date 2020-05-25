package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)

type Fall struct {
	nMenu  *fyne.MainMenu
}

// var aFal = Fall{}
// var Item []*fyne.MenuItem

func main() {
	size := fyne.NewSize(500,200)
	app := app.New()
	w := app.NewWindow("Hello")
	w.Resize(size)
	w.SetMainMenu(createMainMenu())
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}

func createMainMenu() *fyne.MainMenu {
	ItemM := fyne.NewMenuItem("Cours",func(){
		fmt.Println("Ok Ok")
	})
	equation := fyne.NewMenuItem("Equation",func(){
		fmt.Println("Ok Ok")
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
