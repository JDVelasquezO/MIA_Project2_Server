package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"strconv"
)

func PostPrediction (c *fiber.Ctx) error {
	cookie := c.Cookies("user")
	var data map[string]int // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	userRes1 := data["userRes1"]
	userRes2 := data["userRes2"]
	idEvent := data["id_event"]
	idTeam1 := data["id_team1"]
	idTeam2 := data["id_team2"]
	/*idClass1 := data["id_class1"]
	idClass2 := data["id_class2"]*/

	queryGetIdMembership := "SELECT ID_MEMBERSHIP FROM MEMBERSHIP " +
		"WHERE FK_IDUSER = "+cookie+" "
	rows0, err0 := database.DB.Query(queryGetIdMembership)
	if err0 != nil {
		fmt.Println("Error en la consulta 2")
		log.Fatal(err0)
		return err0
	}
	var idMembership int
	for rows0.Next() {
		err := rows0.Scan(&idMembership)
		if err != nil {
			return err
		}
	}
	println(idMembership)

	// Para generar ids
	var countIds int
	rows1, _ := database.DB.Query("SELECT COUNT(*) FROM PREDICTION")
	for rows1.Next() {
		err := rows1.Scan(&countIds)
		if err != nil {
			return err
		}
	}
	// println(countIds)
	firstId := countIds + 1
	secondId := firstId + 1

	// Para insertar predicciones:
	queryInsertPrediction := "INSERT ALL " +
		"INTO PREDICTION (ID_PREDICTION, USER_RESULT, FK_IDMEMBERSHIP) " +
		"VALUES ("+strconv.Itoa(firstId)+", "+strconv.Itoa(userRes1)+", "+strconv.Itoa(idMembership)+") " +
		"INTO PREDICTION (ID_PREDICTION, USER_RESULT, FK_IDMEMBERSHIP) " +
		"VALUES ("+strconv.Itoa(secondId)+", "+strconv.Itoa(userRes2)+", "+strconv.Itoa(idMembership)+") " +
		"SELECT 1 FROM DUAL"
	_, err2 := database.DB.Query(queryInsertPrediction)
	if err2 != nil {
		fmt.Println("Error en la consulta 2")
		log.Fatal(err2)
		return err2
	}

	queryRelationEvent := "INSERT ALL " +
		"INTO PREDICTION_EVENT (FK_IDPREDICTION, FK_IDEVENT, FK_IDTEAM) " +
		"VALUES ("+strconv.Itoa(firstId)+", "+strconv.Itoa(idEvent)+", "+strconv.Itoa(idTeam1)+") " +
		"INTO PREDICTION_EVENT (FK_IDPREDICTION, FK_IDEVENT, FK_IDTEAM) " +
		"VALUES ("+strconv.Itoa(secondId)+", "+strconv.Itoa(idEvent)+", "+strconv.Itoa(idTeam2)+") " +
		"SELECT 1 FROM DUAL"
	_, err3 := database.DB.Query(queryRelationEvent)
	if err3 != nil {
		fmt.Println("Error en la consulta 2")
		log.Fatal(err3)
		return err3
	}

	return c.JSON(fiber.Map{
		"msg": "success",
	})
}
