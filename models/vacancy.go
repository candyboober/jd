package models

type Vacancy struct {
	ID uint `gorm:"primary_key"`
	Title string `json:"title" gorm:"size:80; not null"`
	Body  string `json:"body" gorm:"size:3000; not null"`
}
