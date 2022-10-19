package repository

import(
	"finalproject/app/models"
	"finalproject/config"
)

type PhotoRepository interface{
	CreatePhoto(Photo models.Photo)(models.Photo, error)
	GetPhotobyUserId(user_id *uint)([]models.Photo, error)
	GetPhotobyId(id uint)(models.Photo, error)
	UpdatePhoto(id uint, Photo models.Photo)(models.Photo, error)
	DeletePhoto(id uint)error
	GetAllPhoto()([]models.Photo, error)
}



func NewPhotoRepository() PhotoRepository{
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func(db *dbConnection) CreatePhoto(Photo models.Photo)(models.Photo, error){
	err := db.connection.Save(&Photo).Error
	if err != nil {
		return Photo, err
	}
	return Photo, nil
}

func(db *dbConnection) GetPhotobyUserId(user_id *uint)([]models.Photo, error){
	var Photo []models.Photo
	connection := db.connection.Where("user_id = ?", user_id).Preload("User").Find(&Photo)
	err := connection.Error
	if err != nil{
		return Photo, err
	}
	return Photo, nil
}

func(db *dbConnection) GetPhotobyId(id uint)(models.Photo, error){
	var Photo models.Photo
	connection := db.connection.Where("id = ?", id).Find(&Photo)
	err := connection.Error
	if err != nil{
		return Photo, err
	}
	return Photo, nil
}

func(db *dbConnection) GetAllPhoto()([]models.Photo, error){
	var Photo []models.Photo
	connection := db.connection.Find(&Photo)
	err := connection.Error
	if err != nil{
		return Photo, err
	}
	return Photo, nil
}

func(db *dbConnection) UpdatePhoto(id uint, Photo models.Photo)(models.Photo, error){
	err := db.connection.Model(&Photo).Where("id = ?", id).Updates(&Photo).Error
	if err != nil {
		return Photo, err
	}
	return Photo, nil
}

func (db *dbConnection) DeletePhoto(id uint)error{
	var Photo models.Photo
	err := db.connection.Where("id = ?", id).Delete(&Photo).Error
	if err != nil {
		return err
	}
	return nil
}