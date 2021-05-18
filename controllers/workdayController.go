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

/*func PostWorkday (c *fiber.Ctx) error {
	var data map[string]string // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	var startDate = data["startDate"]
	var endDate = data["endDate"]
	var idSeason = data["idSeason"]
	idSeason += 1
	query2 := "INSERT INTO SEASON (ID_SEASON, START_DATE, END_DATE, NAME) " +
		"VALUES ("+strconv.Itoa(idSeason)+", TO_DATE('"+startDate+"', 'yyyy/mm/dd hh24:mi'), " +
		" TO_DATE('"+endDate+"', 'yyyy/mm/dd hh24:mi'), '"+nameSeason+"' )"
	_, err := database.DB.Query(query2)
	if err != nil {
		return err
	}
}*/
