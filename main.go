package main

import(
	"fmt"
	"finalproject/app/routers"
	"finalproject/config"
	"gorm.io/gorm"
)

var(
	db *gorm.DB = config.ConnectDB()
)

// @title Final Project by Bintang
// @version 1.0
// @description This is a service for final project
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email kadekbintanga@gmail.com
// @license.name Apache 2.0
// @host: localhost:8080
// @BasePath /

func main(){
	fmt.Println("------------------- Final Project Start -------------------")
	config.LoadEnv()
	config.MigrateDatabase(db)
	defer config.DisconnectDB(db)
	routers.InitRouter()
}