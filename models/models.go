package models

import (
	"fmt"
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
	var err error
	this.Connect, err = gorm.Open("postgres", settings)

	if err != nil {
		fmt.Println(settings)
		panic(err)
	}
	this.Connect.LogMode(this.LogMode)
}

// base model
type JDTable struct {
	gorm.Model
	DB DB
}

// create table if not exists
func (this *JDTable) CreateTable(db gorm.DB) {
	if !db.HasTable(this) {
		db.CreateTable(this)
	}
}

// CRUD methods
func (this *JDTable) Get(id string) {
	//this.DB.Connect.Where("id = ?", id).First(this)
	Database.Connect.Where("id = ?", id).First(this)
}

// models
type Vacancy struct {
	JDTable

	Title string
	Body  string `gorm:"size:3000"`
}

func init() {
	databaseSetting := "user=candy dbname=jd password=1 sslmode=disable"
	db, err := gorm.Open("postgres", databaseSetting)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	Database.Init(databaseSetting)
}
