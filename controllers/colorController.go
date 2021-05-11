package controllers

import (
	"github.com/gofiber/fiber/v2"
	"server/database"
	"server/models"
	"strconv"
)

func GetColors (c *fiber.Ctx) error {
	query := "SELECT ID_COLOR, NAME_COLOR, COD_HEX FROM COLOR ORDER BY ID_COLOR"
	rows, err := database.DB.Query(query)
	if err != nil {
		println(err)
		return err
	}

	var colors[]models.Color
	for rows.Next() {
		var color models.Color
		var idColor int
		var nameColor string
		var codHex string
		err := rows.Scan(&idColor, &nameColor, &codHex)
		if err != nil {
			return err
		}

		color.IdColor = idColor
		color.NameColor = nameColor
		color.CodeHex = codHex
		colors = append(colors, color)
	}

	return c.JSON(colors)
}

func PostColor (c *fiber.Ctx) error {
	var data map[string]string // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	var counterColor int
	query := "SELECT COUNT(*) FROM COLOR"
	row, _ := database.DB.Query(query)
	for row.Next() {
		err := row.Scan(&counterColor)
		if err != nil {
			println("Error 1")
			return err
		}
	}

	var idColor = counterColor + 1
	var nameColor = data["nameColor"]
	var codHex = data["codHex"]
	query2 := "INSERT INTO COLOR (ID_COLOR, NAME_COLOR, COD_HEX) " +
		"VALUES ("+strconv.Itoa(idColor)+", '"+nameColor+"', '"+codHex+"') "
	_, err := database.DB.Query(query2)
	if err != nil {
		println("Error")
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "success",
	})
}
