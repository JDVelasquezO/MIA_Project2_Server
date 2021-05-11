package controllers

import (
	"github.com/gofiber/fiber/v2"
	"server/database"
	"server/models"
)

func GetWorkday (c *fiber.Ctx) error {
	query := "SELECT ID_WORKINGDAY, NAME FROM WORKING_DAY"
	rows, err := database.DB.Query(query)
	if err != nil {
		println(err)
		return err
	}

	var workdays []models.Workday
	for rows.Next() {
		var workday models.Workday
		var idWD int
		var nameWD string
		err := rows.Scan(&idWD, &nameWD)
		if err != nil {
			println(err)
			return err
		}

		workday.IdWorkday = idWD
		workday.NameWorkday = nameWD
		workdays = append(workdays, workday)
	}

	return c.JSON(workdays)
}
