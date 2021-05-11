package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/models"
	"strconv"
)

func GetTeamById(c *fiber.Ctx) error {
	var data map[string]int
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	var IdSport = data["idSport"]
	query := "SELECT IDTEAM, NAME_TEAM FROM TEAM " +
		"WHERE FK_IDSPORT = "+strconv.Itoa(IdSport)+" " +
		"ORDER BY IDTEAM"
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal("Error en la consulta")
		return err
	}

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		var idTeam int
		var nameTeam string
		err := rows.Scan(&idTeam, &nameTeam)
		if err != nil {
			return err
		}

		team.IdTeam = idTeam
		team.NameTeam = nameTeam
		teams = append(teams, team)
	}

	return c.JSON(teams)
}
