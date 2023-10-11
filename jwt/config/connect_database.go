package config

import (
	"jwt/mosels"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := "root:klmnopq1@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("errrooror")
	}
	DB.AutoMigrate(&mosels.User{},&mosels.Product{})
}
