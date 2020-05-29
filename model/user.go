package model

import (
	
	"github.com/jinzhu/gorm"
)

type User struct {
	Id string 		`json:"id"`
	UserId int    	`json:"userId"`
	ApiKey  string 	`json:"apiKey"`
	Useremail  string 	`json:"useremail"`
}

func (u *User) FindUser(db *gorm.DB, apikey string) (*User,error) {
	var err error
	var res = User{} 
	err = db.Debug().Model(&User{}).Where("user_api_key = ?", apikey).Find(&User{}).Error

	if err != nil {
		return &User{}, err
	}

	return &res, nil
}

