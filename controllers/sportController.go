package controllers

import (
	"github.com/gofiber/fiber/v2"
	"server/database"
	"server/models"
)

func GetSports(c *fiber.Ctx) error {
	query := "SELECT ID_SPORT, NAME_SPORT FROM SPORT"
	rows, err := database.DB.Query(query)
	if err != nil {
		println(err)
		return err
	}

	var sports []models.Sport
	for rows.Next() {
		var sport models.Sport
		var idSport int
		var nameSport string
		err := rows.Scan(&idSport, &nameSport)
		if err != nil {
			println(err)
			return err
		}
		sport.IdSport = idSport
		sport.NameSport = nameSport
		sports = append(sports, sport)
	}

	return c.JSON(sports)
}
