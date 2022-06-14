package main

import (
	"github.com/gofiber/fiber"
	"github.com/rootturk/goreactjwt/routes"
	"github.com/rootturk/goreactjwt/database"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	database.setupDB()

	app.Listen(":3200")
}
