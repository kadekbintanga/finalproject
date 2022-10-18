package repository

import(
	"finalproject/app/models"
	"finalproject/config"
)

type SocialMediaRepository interface{
	CreateSocialMedia(SocialMedia models.SosialMedia)(models.SosialMedia, error)
	GetSocialMediabyUserId(user_id *uint)([]models.SosialMedia, error)
	UpdateSocialMedia(id uint, SocialMedia models.SosialMedia)(models.SosialMedia, error)
	DeleteSocialMedia(id uint)error
	GetSocialMediabyId(id uint)(models.SosialMedia, error)
}

func NewSocialMediaRepository() SocialMediaRepository{
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func(db *dbConnection) CreateSocialMedia(SocialMedia models.SosialMedia)(models.SosialMedia, error){
	err := db.connection.Save(&SocialMedia).Error
	if err != nil {
		return SocialMedia, err
	}
	return SocialMedia, nil
}

func(db *dbConnection) GetSocialMediabyUserId(user_id *uint)([]models.SosialMedia, error){
	var SocialMedia []models.SosialMedia
	connection := db.connection.Where("user_id= ?", user_id).Preload("User").Find(&SocialMedia)
	err := connection.Error
	if err != nil {
		return SocialMedia, err
	}
	return SocialMedia, nil
}

func(db *dbConnection) GetSocialMediabyId(id uint)(models.SosialMedia, error){
	var SocialMedia models.SosialMedia
	connection := db.connection.Where("id= ?", id).Find(&SocialMedia)
	err := connection.Error
	if err != nil {
		return SocialMedia, err
	}
	return SocialMedia, nil
}

func(db *dbConnection) UpdateSocialMedia(id uint, SocialMedia models.SosialMedia)(models.SosialMedia, error){
	err := db.connection.Model(&SocialMedia).Where("id = ?", id).Updates(&SocialMedia).Error
	if err != nil {
		return SocialMedia, err
	}
	return SocialMedia, nil
}

func (db *dbConnection) DeleteSocialMedia(id uint)error{
	var SocialMedia models.SosialMedia
	err := db.connection.Where("id = ?", id).Delete(&SocialMedia).Error
	if err != nil {
		return err
	}
	return nil
}