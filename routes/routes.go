package routes

import (
	"github.com/gofiber/fiber"
	"github.com/rootturk/goreactjwt/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/login", controllers.Login)
	app.Post("/api/topics", controllers.Login)
}
