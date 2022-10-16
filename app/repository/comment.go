package repository

import(
	"finalproject/app/models"
	"finalproject/config"
)

type CommentRepository interface{
	CreateComment(Comment models.Comment)(models.Comment, error)
	GetCommentbyUserId(user_id uint)([]models.Comment, error)
	GetCommentbyId(id uint)(models.Comment, error)
	UpdateComment(id uint, Comment models.Comment)(models.Comment, error)
	DeleteComment(id uint)error
}

func NewCommentRepository() CommentRepository{
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func(db *dbConnection) CreateComment(Comment models.Comment)(models.Comment, error){
	err := db.connection.Save(&Comment).Error
	if err != nil {
		return Comment, err
	}
	return Comment, nil
}

func(db *dbConnection) GetCommentbyUserId(user_id uint)([]models.Comment, error){
	var Comment []models.Comment
	connection := db.connection.Where("user_id = ?", user_id).Preload("User").Preload("Photo").Find(&Comment)
	err := connection.Error
	if err != nil{
		return Comment, err
	}
	return Comment, nil
}

func(db *dbConnection) GetCommentbyId(id uint)(models.Comment, error){
	var Comment models.Comment
	connection := db.connection.Where("id = ?", id).Find(&Comment)
	err := connection.Error
	if err != nil{
		return Comment, err
	}
	return Comment, nil
}

func(db *dbConnection) UpdateComment(id uint, Comment models.Comment)(models.Comment, error){
	err := db.connection.Model(&Comment).Where("id = ?", id).Updates(&Comment).Error
	if err != nil {
		return Comment, err
	}
	return Comment, nil
}

func (db *dbConnection) DeleteComment(id uint)error{
	var Comment models.Comment
	err := db.connection.Where("id = ?", id).Delete(&Comment).Error
	if err != nil {
		return err
	}
	return nil
}