package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/models"
)

func UploadFile(c *fiber.Ctx) error {
	var data = make(map[string]map[string]models.ModelBL)
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	for d := range data["load"] {
		first := data["load"][d].Firstname
		last := data["load"][d].Lastname
		username := data["load"][d].Username
		pass := data["load"][d].Password
		insertUser(username, pass, first, last, "1999-08-15")
	}

	// return c.JSON(data["load"]["A10"].FootballPool[0].NameSeason)
	return c.JSON(fiber.Map{
		"msg": "success",
	})
}

func insertUser (username string, password string, first string, last string, dateBirth string) {
	println(username)
	query := "INSERT INTO TEST.USERS (USERNAME, PASSWORD, FIRST_NAME, LAST_NAME, DATE_BIRTH," +
		"DATE_REGISTER, PATH_PHOTO, FK_IDROLE) VALUES ('" + username + "', '" + password + "', " +
		"'" + first + "', '" + last + "', TO_DATE('" + dateBirth + "', 'yyyy/mm/dd')," +
		"TO_DATE('" + dateBirth + "', 'yyyy/mm/dd'), 'https://logodix.com/logo/2142984.png', 4)"

	_, err := database.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertSeason (startDate string, endDate string, nameSeasson string) {
	println(startDate, endDate, nameSeasson)
	/*var idSeason int
	query := "SELECT COUNT(*) FROM SEASON"
	rows, _ := database.DB.Query(query)
	for rows.Next() {
		err := rows.Scan(&idSeason)
		if err != nil {
			log.Fatal(err)
		}
	}

	queryInsertSeason := "INSERT INTO SEASON (ID_SEASON, START_DATE, END_DATE, NAME) " +
		"VALUES ("+strconv.Itoa(idSeason)+", TO_DATE("+startDate+", yyyy/mm/dd hh24:mi)), " +
		"TO_DATE("+endDate+", yyyy/mm/dd hh24:mi), '"+nameSeasson+"' "

	_, err := database.DB.Query(queryInsertSeason)
	if err != nil {
		log.Fatal(err)
	}*/
}