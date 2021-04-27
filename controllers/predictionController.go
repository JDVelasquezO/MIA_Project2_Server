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
	// idMembership := data["id_membership"]
	idEvent := data["id_event"]

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

	// OBTENER IDS DE LAS PREDICCIONES
	queryGetIdsPrediction := "SELECT FK_IDPREDICTION FROM EVENT_HAS_TEAM " +
		"INNER JOIN PREDICTION P on EVENT_HAS_TEAM.FK_IDPREDICTION = P.ID_PREDICTION " +
		"WHERE FK_IDEVENT = "+strconv.Itoa(idEvent)+" " +
		"AND FK_IDMEMBERSHIP = "+strconv.Itoa(idMembership)+" "

	rows, err := database.DB.Query(queryGetIdsPrediction)
	if err != nil {
		fmt.Println("Error en la consulta 2")
		log.Fatal(err)
		return err
	}

	var idsPredictions []int
	var idPrediction int
	for rows.Next() {
		err := rows.Scan(&idPrediction)
		if err != nil {
			return err
		}
		idsPredictions = append(idsPredictions, idPrediction)
	}

	// ACTUALIZAR PREDICCIONES
	queryUpdatePrediction1 := "UPDATE PREDICTION SET " +
		"USER_RESULT = "+strconv.Itoa(userRes1)+" " +
		"WHERE ID_PREDICTION = "+strconv.Itoa(idsPredictions[0])+" "
	_, err2 := database.DB.Query(queryUpdatePrediction1)
	if err != nil {
		fmt.Println("Error en la consulta 2")
		log.Fatal(err2)
		return err2
	}

	queryUpdatePrediction2 := "UPDATE PREDICTION SET " +
		"USER_RESULT = "+strconv.Itoa(userRes2)+" " +
		"WHERE ID_PREDICTION = "+strconv.Itoa(idsPredictions[1])+" "
	_, err3 := database.DB.Query(queryUpdatePrediction2)
	if err != nil {
		fmt.Println("Error en la consulta 2")
		log.Fatal(err3)
		return err3
	}

	return c.JSON(fiber.Map{
		"userRes1": userRes1,
		"userRes2": userRes2,
	})
}
