package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"jd/models"
)

type DB struct {
	Connect *gorm.DB
	LogMode bool
	PageSize uint16
}
// initialize database
func (this *DB) Init(settings string) {
	this.LogMode = PostgresLogMode
	this.PageSize = PageSize

	var err error
	this.Connect, err = gorm.Open("postgres", settings)

	if err != nil {
		panic(err)
	}
	this.Connect.LogMode(this.LogMode)
}

func (this *DB) PaginatedQuery(page int) *gorm.DB {
	offset := int(this.PageSize) * page - int(this.PageSize)
	return this.Connect.Offset(offset).Limit(int(this.PageSize))
}

var Database DB

func init() {
	databaseSetting := "user=candy dbname=jd password=1 sslmode=disable"
	Database.Init(databaseSetting)

	Database.Connect.AutoMigrate(&models.Vacancy{})
	Database.Connect.AutoMigrate(&models.User{})
	Database.Connect.AutoMigrate(&models.Message{})
}
