package model

import (
	"fmt"
	"strconv"
	"time"

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
	now := time.Now()
	y := now.Year()
	m := now.Month()
	d := now.Day()
	h := now.Hour()
	min := now.Minute()

	if min < 9 {
		min = fmt.Sprintf("%s%d", "0", min)
	} else {
		min = strconv.Itoa(min)
	}

	last := fmt.Sprintf("%d%s%d%s%d%s%d%s%s%s%s", y, "-", m, "-", d, " ", h, ":", min, ":", "00")
	preview := fmt.Sprintf("%d%s%d%s")
	fmt.Println(last, preview)
	err = db.Debug().Model(&Log{}).Where("date_request BETWEEN ? AND ?", "2020-05-02 00:00:00", "2020-05-02 23:59:59").Order("date_request desc").Limit(10).Find(&posts).Error

	if err != nil {
		return &[]Log{}, err
	}

	if len(posts) > 0 {
		fmt.Println(posts)
	}
	return &posts, nil
}
