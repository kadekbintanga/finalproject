package models

import(
	"time"
	"gorm.io/gorm"
)



type User struct{
	gorm.Model
	ID			*uint		`json:"id" gorm:"primary_key"`
	Username	string		`json:"username" gorm:"index:username_unique_index;unique"`
	Email		string		`json:"email" gorm:"index:email_unique_index;unique"`
	Password	string		`json:"password"`
	Age			uint		`json:"age"`
	CreatedAt	time.Time	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}

type UserToken struct{
	gorm.Model
	ID			*uint		`json:"id" gorm:"primary_key"`
	UserID		*uint		`json:"user_id"`
	User		User		`gorm:"foreignKey:UserID"`
	Token		string		`json:"token"`
	CreatedAt	time.Time	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}