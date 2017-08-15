package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	Connect *gorm.DB
	LogMode bool
}

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

func init() {
	databaseSetting := "user=candy dbname=jd password=1 sslmode=disable"
	Database.Init(databaseSetting)

	Database.Connect.AutoMigrate(&Vacancy{})
	Database.Connect.AutoMigrate(&User{})
}
