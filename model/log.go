package model

import (
	"fmt"
	"log"
	// "reflect"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// map[string]interface{}

type Log struct {
	Id           int    `gorm:"primary_key;auto_increment" json:"id"`
	DateRequest  string `gorm:"not null" json:"date_request"`
	Input        string `gorm:"not null" json:"input"`
	Output       string `gorm:"not null" json:"output"` //Response
	ApiKey       string `gorm:"not null" json:"api_key"`
	IpAdress     string `gorm:"not null" json:"ip_address"`
	NameService  string `gorm:"not null" json:"name_service"`
	ResponseTime string `gorm:"not null" json:"response_time"`
	UserAgent    string `gorm:"null" json:"user_agent"`
	Status       bool   `gorm:"not null" json:"status"`
}

type Users struct {
	Lastname  string `gorm:"not null" json:"lastname"`
	Firstname string `gorm:"not null" json:"firstname"`
}

type Response struct {
	Uid                     string `gorm:"not null" json:"uid"`
	TypeEnveloppe           string `gorm:"not null" json:"type_enveloppe"`
	Enveloppe               string `gorm:"not null" json:"enveloppe"`
	AdresseExpedition       string `gorm:"not null" json:"adresse_expedition"`
	AdresseDestination      string `gorm:"not null" json:"adresse_destination"`
	Fichier                 string `gorm:"not null" json:"fichier"`
	FichierPrevisualisation string `gorm:"not null" json:"fichier_previsualisation"`
	Variables               string `gorm:"not null" json:"variables"`
	FichierAnnexes          string `gorm:"not null" json:"fichier_annexes"`
}

// Get last Order 20 items every five
func (Log *Log) TableName() string {
	return "log"
}

func (l *Log) FindAllPosts(db *gorm.DB) (*[]Log, error) {
	var err error
	posts := []Log{}
	loc, err := time.LoadLocation("Europe/Paris")

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

	inf := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	for i := 0; i < len(inf); i++ {
		if inf[i] == m.String() && i < 10 {
			aMois = fmt.Sprintf("%s%d", "0", i+1)
		} else if inf[i] == m.String() && i > 9 {
			aMois = strconv.Itoa(i + 1)
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
	preview := fmt.Sprintf("%d%s%s%s%s%s%d%s%s%s%s", y, "-", aMois, "-", aDay, " ", h-1, ":", aMin, ":", "00")
	err = db.Debug().Model(&Log{}).Where("date_request BETWEEN ? AND ? AND name_service <> ? AND name_service <> ?", preview, last, "tracking", "letter_pricing").Order("date_request desc").Limit(10).Find(&posts).Error

	if err != nil {
		return &[]Log{}, err
	}

	return &posts, nil
}

func postsHandler(posts []Log) {
	fmt.Println("pp5", posts[0].Output)
}
