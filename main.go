package main

import (
	"fmt"
	"log"

	"github.com/antonio-nirina/monitore-sp-api/config"
	"github.com/antonio-nirina/monitore-sp-api/model"
)

type StartDate struct {
    Year     int `json:"year"`
    Month    int `json:"month"`
    Day      int `json:"day"`
    Hour     int `json:"hour"`
    Minute   int `json:"minute"`
    Second   int `json:"second"`
}

type Data struct {
	X StartDate
	Y bool
}

var logPost = model.Log{}
var graph []Data
var absc = make(map[string]interface{})
var array []interface{}

func main() {
	err,process := config.Connected()

	if err != nil {
		log.Fatal("database not Connected")
	}

	fmt.Println("database is Connected success")
	logs, err := logPost.FindAllPosts(process.DB)

	if err != nil {
		log.Fatal("error request")
	}

	for _,val := range *logs{
		absc["date"] = val.DateRequest
		absc["out"] = val.Output
		absc["status"] = val.Status
		array = append(array,absc)
	}

	fmt.Println(array)
}
