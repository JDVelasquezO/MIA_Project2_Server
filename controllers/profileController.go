package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"time"
)

func UpdateDataUser (c *fiber.Ctx) error {
	cookie := c.Cookies("user")
	var data map[string]string // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	username := data["username"]
	first := data["first"]
	last := data["last"]
	birth, _ := time.Parse(time.RFC3339, data["birth"])
	birthFormat := birth.Format("2006-01-02")
	email := data["email"]
	// println(birthFormat)

	query := "UPDATE USERS SET USERNAME = '"+username+"', FIRST_NAME = '"+first+"', LAST_NAME = '"+last+"', " +
		"DATE_BIRTH = TO_DATE('"+birthFormat+"', 'yyyy-mm-dd'), EMAIL = '"+email+"' " +
		"WHERE ID_USER = "+cookie+" "

	_, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Error en la consulta")
		log.Fatal(err)
		return err
	}

	return User(c)
}

func UpdatePassword (c *fiber.Ctx) error {
	cookie := c.Cookies("user")
	var data map[string]string // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	oldPass := data["oldPass"]
	newPass := data["newPass"]

	getActualPass := "SELECT PASSWORD FROM USERS WHERE ID_USER = "+cookie+" "
	rows, err := database.DB.Query(getActualPass)
	if err != nil {
		fmt.Println("Error en la consulta 1")
		log.Fatal(err)
		return err
	}

	var actualPass string
	for rows.Next() {
		err := rows.Scan(&actualPass)
		if err != nil {
			return err
		}
	}

	if actualPass != oldPass {
		return c.JSON("not success")
	}

	updatePass := "UPDATE USERS SET PASSWORD = '"+newPass+"' " +
		"WHERE ID_USER = "+cookie+" "
	_, err2 := database.DB.Query(updatePass)
	if err2 != nil {
		fmt.Println("Error en la consulta 2")
		log.Fatal(err2)
		return err2
	}

	return User(c)
}