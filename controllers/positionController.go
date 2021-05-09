package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/models"
	"strconv"
)

func GetP10(c *fiber.Ctx) error {
	cookie := c.Cookies("user")
	id, _ := strconv.Atoi(cookie)

	rows, err := database.DB.Query("SELECT ID_USER, FIRST_NAME, LAST_NAME, COUNT(*) Cantidad_Aciertos " +
		"FROM EVENT_HAS_TEAM " +
		"INNER JOIN PREDICTION P on P.ID_PREDICTION = EVENT_HAS_TEAM.FK_IDPREDICTION " +
		"INNER JOIN MEMBERSHIP M on M.ID_MEMBERSHIP = P.FK_IDMEMBERSHIP " +
		"INNER JOIN USERS U on U.ID_USER = M.FK_IDUSER\nWHERE REAL_RESULT = USER_RESULT " +
		"GROUP BY ID_USER, FIRST_NAME, LAST_NAME ")
	if err != nil {
		fmt.Println("Error en la consulta")
		log.Fatal(err)
		return err
	}

	var puntuations []models.P10
	var p10 models.P10
	var idUser int
	var firsName string
	var lastName string
	var quantHits int
	for rows.Next() {
		err := rows.Scan(&idUser, &firsName, &lastName, &quantHits)
		if err != nil {
			println("Error")
			return err
		}
		p10.IdUser = idUser
		p10.FirstName = firsName
		p10.LastName = lastName
		p10.QuantHits = quantHits / 2

		if idUser == id {
			p10.IsUser = true
		} else {
			p10.IsUser = false
		}

		puntuations = append(puntuations, p10)
	}

	return c.JSON(puntuations)
}
