package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	Connect *gorm.DB
	LogMode bool
}

func (this *DB) Init(settings string) {
	var err error
	this.Connect, err = gorm.Open("postgres", settings)

	if err != nil {
		fmt.Println(settings)
		panic(err)
	}
	this.Connect.LogMode(this.LogMode)
}

type JDTable struct{}

func (this *JDTable) CreateTable(db gorm.DB) {
	if !db.HasTable(this) {
		db.CreateTable(this)
	}
}

type Vacancy struct {
	gorm.Model
	JDTable

	Title string
	Body  string `gorm:"size:3000"`
}

func init() {
	// TODO: move to settings file
	//databaseSetting := "user=candy dbname=jd password=1"
	//db, err := gorm.Open("postgres", databaseSetting)
	//fmt.Println(db)
	//defer db.Close()
	//
	//if err != nil {
	//	panic(err)
	//}
}
