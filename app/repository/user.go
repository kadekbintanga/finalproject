package repository

import(
	"finalproject/app/models"
	"finalproject/config"
)


type UserRepository interface{
	CreateUser(User models.User)(models.User, error)
	GetUserByEmail(email string)(models.User, error)
	GetToken(user_id uint)(models.UserToken, error)
	AddToken(UserToken models.UserToken)(models.UserToken, error)
	DeleteToken(id uint)error
}

func NewUserRepository() UserRepository{
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func (db *dbConnection) CreateUser(User models.User)(models.User, error){
	err := db.connection.Save(&User).Error
	if err != nil {
		return User, err
	}
	return User, nil
}

func (db *dbConnection) GetUserByEmail(email string)(models.User, error){
	var User models.User
	connection := db.connection.Where("email = ?", email).Find(&User)
	err := connection.Error
	if err != nil {
		return User, err
	}
	return User, nil
}

func (db *dbConnection) GetToken(user_id uint)(models.UserToken, error){
	var UserToken models.UserToken
	connection := db.connection.Where("user_id = ?", user_id).Find(&UserToken)
	err := connection.Error
	if err != nil {
		return UserToken, err
	}
	return UserToken, nil
}

func (db *dbConnection) AddToken(UserToken models.UserToken)(models.UserToken, error){
	err := db.connection.Save(&UserToken).Error
	if err != nil {
		return UserToken, err
	}
	return UserToken, nil
}

func (db *dbConnection) DeleteToken(id uint)error{
	var UserToken models.UserToken
	err := db.connection.Delete(&UserToken, id).Error
	if err != nil {
		return err
	}
	return nil
}