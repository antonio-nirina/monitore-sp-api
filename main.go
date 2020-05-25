package main

import (
	"fmt"
	"log"
	"time"
	//"encoding/json"
	//"reflect"

	"github.com/antonio-nirina/monitore-sp-api/config"
	"github.com/antonio-nirina/monitore-sp-api/model"
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

func mainnn() {
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

			if !val.Status {
				graph = append(graph, 0)
			} else {
				graph = append(graph, 1)
			}

			array = append(array, absc)
		}
		fmt.Println(array)
		fmt.Println(graph)
	}

	fmt.Println("array")
}
