package models

import(
	"time"
	"gorm.io/gorm"
)

type SosialMedia struct{
	gorm.Model
	ID					*uint		`json:"id" gorm:"primary_key"`
	Name				string		`json:"name"`
	SocialMediaUrl		string		`json:"social_media_url"`
	UserID				*uint		`json:"user_id"`
	User				User		`gorm:"foreignKey:UserID"`
	CreatedAt			time.Time	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt			time.Time	`json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}