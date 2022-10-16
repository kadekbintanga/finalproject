package repository

import(
	"finalproject/app/models"
	"finalproject/config"
)

type PhotoRepository interface{
	CreatePhoto(Photo models.Photo)(models.Photo, error)
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