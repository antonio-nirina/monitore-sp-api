package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int    `gorm:"primary_key;auto_increment" json:"id"`
	UserId    int    `gorm:"column:user_id" json:"user_id"`
	ApiKey    string `gorm:"column:user_api_key;not null" json:"user_api_key"`
	Useremail string `gorm:"column:user_email;not null" json:"user_email"`
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
	// fmt.Println("aaaa", u.Useremail)
	return u, nil
}
