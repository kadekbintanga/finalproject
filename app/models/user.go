package models

import(
	"time"
	"gorm.io/gorm"
)



type User struct{
	gorm.Model
	ID			*uint		`json:"id" gorm:"primary_key"`
	Username	string		`json:"username"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	Age			string		`json:"age"`
	CreatedAt	time.Time	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}