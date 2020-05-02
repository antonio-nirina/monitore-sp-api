package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// map[string]interface{}

type Log struct {
	Id int `gorm:"primary_key;auto_increment" json:"id"`
	DateRequest string `gorm:"not null" json:"dateRequest"`
	Input string  `gorm:"not null" json:"input"`
	Output string  `gorm:"not null" json:"output"`
	ApiKey string `gorm:"not null" json:"apiKey"`
	IpAdress string `gorm:"not null" json:"ipAdress"`
	NameService string `gorm:"not null" json:"nameService"`
	UserAgent string `gorm:"null" json:"userAgent"`
}

// Get last Order 20 items every five

func (l *Log) FindAllPosts(db *gorm.DB) (*[]Log, error) {
	var err error
	posts := []Log{}
	err = db.Debug().Model(&Log{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Log{}, err
	}

	if len(posts) > 0 {
		fmt.Println(posts)
	}
	return &posts, nil
}