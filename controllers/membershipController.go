package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/models"
	"strconv"
)

func GetStatusMembership(c *fiber.Ctx) error {
	cookie := c.Cookies("user")
	id, _ := strconv.Atoi(cookie)

	query := "SELECT ID_MEMBERSHIP, TIER_NAME, TIER_PRICE FROM MEMBERSHIP " +
		"INNER JOIN TIER T on T.IDTIER = MEMBERSHIP.FK_IDTIER " +
		"INNER JOIN USERS U on U.ID_USER = MEMBERSHIP.FK_IDUSER " +
		"WHERE FK_IDUSER = '" + cookie + "' "
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Error en la consulta")
		log.Fatal(err)
		return err
	}

	var membership models.Membership
	var idMembership int
	var typeTier string
	var priceTier float32
	for rows.Next() {
		err := rows.Scan(&idMembership, &typeTier, &priceTier)
		if err != nil {
			return err
		}
	}

	membership.IdMembership = idMembership
	membership.IdUser = id
	membership.PriceTier = priceTier
	membership.TypeTier = typeTier

	return c.JSON(membership)
}

func UpdateMembership (c *fiber.Ctx) error {
	var data map[string]string // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	IdTier, _ := data["idTier"]
	IdUser, _ := data["idUser"]

	query := "UPDATE MEMBERSHIP SET FK_IDTIER = "+IdTier+" WHERE FK_IDUSER = "+IdUser+" "
	_, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Error en la consulta")
		log.Fatal(err)
		return err
	}

	return GetStatusMembership(c)
}
