package routes

import (
	"github.com/gofiber/fiber/v2"
	"server/controllers"
)

func Setup(app *fiber.App) {
	// Routes of auth
	app.Post("/quinielas.io/register", controllers.Register)
	app.Post("/quinielas.io/login", controllers.Login)
	app.Get("/quinielas.io/user", controllers.User)
	app.Post("/quinielas.io/logout", controllers.Logout)

	// Routes of auth
	app.Get("/quinielas.io/membership", controllers.GetStatusMembership)
	app.Put("/quinielas.io/updateMembership", controllers.UpdateMembership)

	// Routes of profile
	app.Put("/quinielas.io/updateUser", controllers.UpdateDataUser)
	app.Put("/quinielas.io/updatePass", controllers.UpdatePassword)
	app.Put("/quinielas.io/updatePhoto", controllers.UpdateProfilePhoto)

	// Routes of events
	app.Get("/quinielas.io/getEvents", controllers.GetEvents)
	app.Get("/quinielas.io/getEvent/:id", controllers.GetEvent)
	app.Post("/quinielas.io/postEvent", controllers.PostEvent)

	// Routes of predictions
	app.Post("/quinielas.io/postPrediction", controllers.PostPrediction)

	// Routes of seasons
	app.Get("/quinielas.io/getActualSeason", controllers.GetActualSeason)

	// Routes of positions
	app.Get("/quinielas.io/getPositionsP10", controllers.GetP10)

	// Routes for bulk load
	app.Post("/quinielas.io/uploadFile", controllers.UploadFile)

	// Routes for sports
	app.Get("/quinielas.io/getSports", controllers.GetSports)

	// Routes for teams
	app.Post("/quinielas.io/getTeamsById", controllers.GetTeamById)

	// Routes for workdays
	app.Get("/quinielas.io/getWorkdays", controllers.GetWorkday)
}
