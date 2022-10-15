package config

import(
	"finalproject/app/models"
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB{
	LoadEnv()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil{
		panic("Failed to connect postgres database")
	}

	return db
}


func MigrateDatabase(db *gorm.DB){
	db.AutoMigrate(
		&models.User{},&models.UserToken{},
	)

	fmt.Println("Database Migration has been success")
}

func DisconnectDB(db *gorm.DB){
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}