package controllers

import (
	"github.com/gofiber/fiber/v2"
	"server/database"
	"server/models"
	"strconv"
)

func GetSports(c *fiber.Ctx) error {
	query := "SELECT ID_SPORT, NAME_SPORT, COD_HEX " +
		"FROM SPORT " +
		"INNER JOIN COLOR C2 on C2.ID_COLOR = SPORT.FK_IDCOLOR " +
		"ORDER BY ID_SPORT"
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
		var nameColor string
		err := rows.Scan(&idSport, &nameSport, &nameColor)
		if err != nil {
			println(err)
			return err
		}
		sport.IdSport = idSport
		sport.NameSport = nameSport
		sport.NameColor = nameColor
		sports = append(sports, sport)
	}

	return c.JSON(sports)
}

func PostSport (c *fiber.Ctx) error {
	var data map[string]string // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}
	var countSports int
	query := "SELECT COUNT(*) FROM SPORT"
	rows, _ := database.DB.Query(query)
	for rows.Next() {
		err := rows.Scan(&countSports)
		if err != nil {
			println("Error 1")
			return err
		}
	}

	var idSport = countSports + 1
	var nameSport = data["nameSport"]
	var fkIdColor = data["fkIdColor"]

	query2 := "INSERT INTO SPORT (ID_SPORT, NAME_SPORT, FK_IDCOLOR) " +
		"VALUES ("+strconv.Itoa(idSport)+", '"+nameSport+"', "+fkIdColor+") "

	_, err := database.DB.Query(query2)
	if err != nil {
		println("Error 2")
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "success",
	})
}
