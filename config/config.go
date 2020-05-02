package config

import (
	"os"
	"fmt"
	"log"

	"github.com/antonio-nirina/monitore-sp-api/model"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

type Process struct {
	DB *gorm.DB
}


func Connected()(error, Process) {
	err := godotenv.Load()
	process := Process{}

	if err != nil {
		log.Fatal("Error loading .env file",err)
		return err,process
	}

	dbName := os.Getenv("DATABASE")
	dbPwd := os.Getenv("DB_PWD")
	dbUser := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")

	_uri := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPwd, host, dbName)
	process.DB, err = gorm.Open("mysql", _uri)

	if err != nil {
		return err,process
	} 

	process.DB.AutoMigrate(&model.Log{})

	return nil, process
}