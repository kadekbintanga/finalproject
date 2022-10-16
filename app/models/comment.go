package models

import(
	"time"
	"gorm.io/gorm"
)

type Comment struct{
	gorm.Model
	ID			*uint		`json:"id" gorm:"primary_key"`
	Message		string		`json:"message"`
	UserID		*uint		`json:"user_id"`
	User		User		`gorm:"foreignKey:UserID"`
	PhotoID		*uint		`json:"photo_id"`
	Photo		Photo		`gorm:"foreignKey:PhotoID"`
	CreatedAt	time.Time	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}