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
	idClass1 := data["id_class1"]
	idClass2 := data["id_class2"]

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
	idPrediction := (countIds + 2)/2

	// Para insertar predicciones:
	queryInsertPrediction := "INSERT ALL " +
		"INTO PREDICTION (ID_PREDICTION, USER_RESULT, FK_IDMEMBERSHIP, FK_IDEVENT, DATE_REGISTER, FK_IDCLASS) " +
		"VALUES ("+strconv.Itoa(idPrediction)+", "+strconv.Itoa(userRes1)+", "+strconv.Itoa(idMembership)+", " +
		" "+strconv.Itoa(idEvent)+", SYSDATE, "+strconv.Itoa(idClass1)+") " +
		"INTO PREDICTION (ID_PREDICTION, USER_RESULT, FK_IDMEMBERSHIP, FK_IDEVENT, DATE_REGISTER, FK_IDCLASS) " +
		"VALUES ("+strconv.Itoa(idPrediction)+", "+strconv.Itoa(userRes2)+", "+strconv.Itoa(idMembership)+", " +
		" "+strconv.Itoa(idEvent)+", SYSDATE, "+strconv.Itoa(idClass2)+" ) " +
		"SELECT 1 FROM DUAL"
	_, err2 := database.DB.Query(queryInsertPrediction)
	if err2 != nil {
		fmt.Println("Error en la consulta 2")
		log.Fatal(err2)
		return err2
	}

	return c.JSON(fiber.Map{
		"msg": "success",
	})
}
