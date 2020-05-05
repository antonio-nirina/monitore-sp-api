package model

import (
	"fmt"
	"strconv"
	"time"
	"log"
	"reflect"

	"github.com/jinzhu/gorm"
)

// map[string]interface{}

type Log struct {
	Id           int    `gorm:"primary_key;auto_increment" json:"id"`
	DateRequest  string `gorm:"not null" json:"date_request"`
	Input        string `gorm:"not null" json:"input"`
	Output       string `gorm:"not null" json:"output"`
	ApiKey       string `gorm:"not null" json:"api_key"`
	IpAdress     string `gorm:"not null" json:"ip_address"`
	NameService  string `gorm:"not null" json:"name_service"`
	ResponseTime string `gorm:"not null" json:"response_time"`
	UserAgent    string `gorm:"null" json:"user_agent"`
}

// Get last Order 20 items every five
func (Log *Log) TableName() string {
	return "log"
}

func (l *Log) FindAllPosts(db *gorm.DB) (*[]Log, error) {
	var err error
	posts := []Log{}
	loc,err := time.LoadLocation("Europe/Paris")

	if err != nil {
		log.Fatal("error location")
	}
	now := time.Now().In(loc)
	y := now.Year()
	m := now.Month()
	d := now.Day()
	h := now.Hour()
	min := now.Minute()
	var aMin string
	var aMois string
	var aDay string
	
	inf := [12]string{"January","February","March","April","May","June","July","August","September","October","November","December"}

	for i := 0; i < len(inf); i++ {
		if inf[i] == m.String() && i < 10{
			aMois = fmt.Sprintf("%s%d", "0", i+1)
		} else if(inf[i] == m.String() && i > 9) {
			aMois = strconv.Itoa(i+1)
		}
	} 

	if d < 10 {
		aDay = fmt.Sprintf("%s%s", "0", strconv.Itoa(d))
	} else {
		aDay = strconv.Itoa(d)
	}

	if min < 10 {
		aMin = fmt.Sprintf("%s%s", "0", strconv.Itoa(min))
	} else {
		aMin = strconv.Itoa(min)
	}
	
	last := fmt.Sprintf("%d%s%s%s%s%s%d%s%s%s%s", y, "-", aMois, "-", aDay, " ", h, ":", aMin, ":", "00")
	preview := fmt.Sprintf("%d%s%s%s%s%s%d%s%s%s%s", y, "-", aMois, "-", aDay, " ", h - 1, ":", aMin, ":", "00")
	err = db.Debug().Model(&Log{}).Where("date_request BETWEEN ? AND ?", preview, last).Order("date_request desc").Limit(10).Find(&posts).Error

	if err != nil {
		return &[]Log{}, err
	}

	if len(posts) > 0 {
		fmt.Println(reflect.TypeOf(posts))
	}
	return &posts, nil
}
