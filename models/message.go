package models

import "time"

type Message struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Author      User      `json:"author" gorm:"ForeignKey:AuthorID"`
	AuthorID    uint      `json:"author_id" gorm:"not null"`
	Recipient   User      `json:"recipient" gorm:"ForeignKey:RecepientID"`
	RecipientID uint      `json:"recipient_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" sql:"DEFAULT:current_timestamp"`
	Body        string    `json:"body" gorm:"not null; size:3000"`
}
