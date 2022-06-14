package database

import (
	"github.com/rootturk/goreactjwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func setupDB() {
	dsn := "root:example@tcp(127.0.0.1:3306)/eskiyen?charset=utf8mb4&parseTime=True&loc=Local"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database Error")
	}

	DB = connection

	connection.AutoMigrate(&models.Topics{}, &models.Login{}, models.Replies{}, models.BlackList{})
}
