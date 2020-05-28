package main

import (
	"fmt"
	"log"
	"time"

	//"encoding/json"
	//"reflect"
	"github.com/antonio-nirina/monitore-sp-api/config"
	"github.com/antonio-nirina/monitore-sp-api/model"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type StartDate struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

type Data struct {
	X StartDate
	Y bool
}

var logPost = model.Log{}
var graph []int
var absc = make(map[string]interface{})
var array []interface{}

func main() {
	err, process := config.Connected()

	if err != nil {
		log.Fatal("database not Connected")
	}

	fmt.Println("database is Connected success")
	logs, err := logPost.FindAllPosts(process.DB)

	if err != nil {
		log.Fatal("error request")
	}

	if len(*logs) > 0 {
		for _, val := range *logs {
			t, _ := time.Parse(time.RFC3339, val.DateRequest)
			absc["id"] = val.Id
			absc["date"] = t.Format("20060102150405") // timestamp Go
			absc["apiKey"] = val.ApiKey
			absc["nameService"] = val.NameService
			absc["status"] = val.Status

			if !val.Status {
				graph = append(graph, 0)
			} else {
				i := 1
				graph = append(graph, float64(i))
			}

			array = append(array, absc)
		}
		traceLogs(graph)
	}

	fmt.Println("array")
}

func traceLogs(data []float64) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	Data := func() [][]int {
		obj := make([][]int, 2)
		obj[0] = make([]int, 2)
		obj[1] = make([]int, 2)
		for _, v := range data {
			obj[0] = append(obj[0], v)
		}
		obj[0] = append(obj[0], 1, 1, 1, 1, 1, 1, 1)
		return obj
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
	p0.Data = testData
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
