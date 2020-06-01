package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int    `gorm:"primary_key;auto_increment" json:"id"`
	UserId    int    `gorm:"not null" json:"user_id"`
	ApiKey    string `gorm:"not null" json:"user_api_key"`
	Useremail string `gorm:"not null" json:"user_email"`
}

func (user *User) TableName() string {
	return "user_api_key"
}

func (u *User) FindUser(db *gorm.DB, apikey string) (*User, error) {
	var err error
	// var res = User{}
	err = db.Debug().Model(&User{}).Where("user_api_key = ?", apikey).Find(&u).Error

	if err != nil {
		return &User{}, err
	}
	fmt.Println("aaaa", u)
	return u, nil
}
