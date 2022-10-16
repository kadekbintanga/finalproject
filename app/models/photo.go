package models

import(
	"time"
	"gorm.io/gorm"
)

type Photo struct{
	gorm.Model
	ID			*uint		`json:"id" gorm:"primary_key"`
	Title		string		`json:"title"`
	Caption		string		`json:"caption"`
	PhotoUrl	string		`json:"photo_url"`
	UserID		*uint		`json:"user_id"`
	User		User		`gorm:"foreignKey:UserID"`
	CreatedAt	time.Time	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}