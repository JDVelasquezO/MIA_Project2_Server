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

func PostTeam(c *fiber.Ctx) error {
	var data map[string]string // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	var counterTeam int
	query := "SELECT COUNT(*) FROM TEAM"
	row, _ := database.DB.Query(query)
	for row.Next() {
		err := row.Scan(&counterTeam)
		if err != nil {
			println("Error 1")
			return err
		}
	}

	var idTeam = counterTeam + 1
	var nameTeam = data["nameTeam"]
	var fkIdSport = data["idSport"]
	query2 := "INSERT INTO TEAM (IDTEAM, NAME_TEAM, FK_IDSPORT) " +
		"VALUES ("+strconv.Itoa(idTeam)+", '"+nameTeam+"', "+fkIdSport+") "
	_, err := database.DB.Query(query2)
	if err != nil {
		println("Error")
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "success",
	})
}
