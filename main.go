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

func main(){
	fmt.Println("------------------- Final Project Start -------------------")
	config.LoadEnv()
	config.MigrateDatabase(db)
	defer config.DisconnectDB(db)
	routers.InitRouter()
}