package main

import (
	"github.com/gofiber/fiber"
	"github.com/rootturk/goreactjwt/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	initilazeDatabase()

	app.Listen(":3200")
}

func initilazeDatabase() {
	dsn := "root:example@tcp(127.0.0.1:3306)/birdsources?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Error")
	}
}
