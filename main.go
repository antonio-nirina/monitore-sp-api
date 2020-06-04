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

type DataResp struct {
	Id          int
	Apikey      string
	Email       string
	StatusLog   string
	Date        string
	NameService string
	StatusInit  bool
}

var logPost = model.Log{}
var graph []float64
var absc = make(map[string]interface{})
var array []DataResp

func checkConnected() config.Process {
	err, process := config.Connected()

	if err != nil {
		log.Fatal("database not Connected")
	}
	fmt.Println("database is Connected success")

	return process
}
func main() {
	process := checkConnected()
	// var k int
	var resp = DataResp{}
	// for k == 3 {
	logs, err := logPost.FindAllPosts(process.DB)

	if err != nil {
		log.Fatal("error request")
	}

	if len(*logs) > 0 {
		var i int
		var st string
		for _, val := range *logs {
			t, _ := time.Parse(time.RFC3339, val.DateRequest)

			if !val.Status {
				i = 0
				st = "Error"
				graph = append(graph, float64(i))
			} else {
				i = 1
				st = "Success"
				graph = append(graph, float64(i))
			}
			resp.Id = val.Id
			resp.Apikey = val.ApiKey
			resp.StatusLog = st
			resp.NameService = val.NameService
			resp.StatusInit = val.Status
			resp.Date = t.Format("20060102150405")

			array = append(array, resp)
		}
		traceLogs(graph, array)
		// time.Sleep(600)
	}

	//fmt.Println(k)
	// k++
	// }
}

func getNameByApikey(apikey string) *model.User {
	process := checkConnected()
	var userSend = model.User{}
	user, err := userSend.FindUser(process.DB, apikey)

	if err != nil {
		log.Fatal("error_find_user")
	}

	return user
}

func traceLogs(data []float64, array []DataResp) {
	if err := ui.Init(); err != nil {
	}
	defer ui.Close()

	Data := func() [][]float64 {
		obj := make([][]float64, 2)
		obj[0] = make([]float64, 2)
		obj[1] = make([]float64, 2)
		for _, v := range data {
			obj[0] = append(obj[0], v)
		}

		return obj
	}()

	// Stat error if exist previsualisation :
	// Nombres: 20
	// Error 500: 0 ou 2
	// Client: Harmonie
	var statutLog string

	for _, val := range array {
		if !val.StatusInit {
			statutLog = "Error"
			user := getNameByApikey(val.Apikey)
			val.Email = user.Useremail
		} else {
			statutLog = "Success"
		}

		p := widgets.NewParagraph()
		d := widgets.NewParagraph()
		em := widgets.NewParagraph()
		serv := widgets.NewParagraph()
		d.Text = "Date: " + val.Date
		em.Text = "Email: " + val.Email
		serv.Text = "Service: " + val.NameService
		p.Text = statutLog
		p.SetRect(3, 0, 12, 3)
		ui.Render(p)
		d.SetRect(3, 5, 12, 3)
		ui.Render(d)
		em.SetRect(3, 2, 12, 3)
		ui.Render(em)
		serv.SetRect(3, 2, 12, 3)
		ui.Render(serv)
	}

	// Plot service Direct and Previsualisation
	p0 := widgets.NewPlot()
	p0.Title = "Monitore sp-api"
	p0.Data = Data
	p0.SetRect(19, 0, 60, 15)
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
