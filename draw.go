package main

import (
	"log"
	"math"
	//"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func maindd() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	sinData := func() [][]float64 {
		n := 20
		data := make([][]float64, 2)
		data[0] = make([]float64, n)
		data[1] = make([]float64, n)
		for i := 0; i < n; i++ {
			data[0][i] = 1 + math.Sin(float64(i)/5)
			//data[1][i] = 1 + math.Cos(float64(i)/5)
		}
		return data
	}()
	// Stat error if exist previsualisation : 
	// Nombres: 20
	// Error 500: 0 ou 2
	// Client: Harmonie
	p := widgets.NewParagraph()
    p.Text = "Hello World!"
    p.SetRect(3, 0, 12, 15)
    ui.Render(p)
    // Plot service Direct and Previsualisation
	p0 := widgets.NewPlot()
	p0.Title = "braille-mode Line Chart"
	p0.Data = sinData
	p0.SetRect(13, 0, 60, 15)
	p0.AxesColor = ui.ColorWhite
	p0.LineColors[0] = ui.ColorRed
	ui.Render(p0)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
