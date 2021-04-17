package routes

import (
	"github.com/gofiber/fiber/v2"
	"server/controllers"
)

func Setup(app *fiber.App)  {
	app.Post("/quinielas.io/register", controllers.Register)
	app.Post("/quinielas.io/login", controllers.Login)
	app.Get("/quinielas.io/user", controllers.User)
	app.Post("/quinielas.io/logout", controllers.Logout)
}
