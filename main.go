package main

import (
	"fmt"
	"log"

	"github.com/antonio-nirina/monitore-sp-api/config"
	"github.com/antonio-nirina/monitore-sp-api/model"
)

var logPost = model.Log{}

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

	fmt.Println(logs)
}
