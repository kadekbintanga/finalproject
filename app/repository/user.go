package repository

import(
	"finalproject/app/models"
	"finalproject/config"
)


type UserRepository interface{
	CreateUser(User models.User)(models.User, error)
	GetUserByEmail(email string)(models.User, error)
	GetToken(user_id *uint)(models.UserToken, error)
	AddToken(UserToken models.UserToken)(models.UserToken, error)
	DeleteToken(user_id *uint)error
	UpdateUser(id *uint, User models.User)(models.User, error)
	DeleteUser(email string)error
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

func (db *dbConnection) GetToken(user_id *uint)(models.UserToken, error){
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

func (db *dbConnection) DeleteToken(user_id *uint)error{
	var UserToken models.UserToken
	err := db.connection.Where("user_id= ?", user_id).Delete(&UserToken).Error
	if err != nil {
		return err
	}
	return nil
}

func(db *dbConnection) UpdateUser(id *uint, User models.User)(models.User, error){
	err := db.connection.Model(&User).Where("id = ?", id).Updates(&User).Error
	if err != nil {
		return User, err
	}
	return User, nil
}

func (db *dbConnection) DeleteUser(email string)error{
	var  User models.User
	err := db.connection.Where("email= ?", email).Delete(&User).Error
	if err != nil {
		return err
	}
	return nil
}