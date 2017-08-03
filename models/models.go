package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// database connection
type DB struct {
	Connect *gorm.DB
	LogMode bool
}

// global objects

// database cursor
var Database DB

// initialize database
func (this *DB) Init(settings string) {
	this.LogMode = true

	var err error
	this.Connect, err = gorm.Open("postgres", settings)

	if err != nil {
		panic(err)
	}
	this.Connect.LogMode(this.LogMode)
}

// CRUD methods
func (this *DB) Get(instance interface{}, id int) {
	this.Connect.First(instance, id)
}

// models
type Vacancy struct {
	ID uint `gorm:"primary_key"`

	Title string `json:"title" gorm:"size:80; not null"`
	Body  string `json: "body" gorm:"size:3000; not null"`
}


func init() {
	databaseSetting := "user=candy dbname=jd password=1 sslmode=disable"
	Database.Init(databaseSetting)

	Database.Connect.AutoMigrate(&Vacancy{})
}
