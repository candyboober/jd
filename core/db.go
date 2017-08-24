package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"jd/models"
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
	this.Connect, err = gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=jd sslmode=disable")
	if err != nil {
		panic(err)
	}
	this.Connect.LogMode(this.LogMode)
}

func init() {
	databaseSetting := "user=candy dbname=jd password=1 sslmode=disable"
	Database.Init(databaseSetting)

	Database.Connect.AutoMigrate(&models.Vacancy{})
	Database.Connect.AutoMigrate(&models.User{})
	Database.Connect.AutoMigrate(&models.Message{})
}
