package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/models"
	"strconv"
	"time"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	var user models.User
	user.Email = data["email"]
	user.Password = data["pass"]
	rows, err := database.DB.Query("SELECT ID_USER, USERNAME " +
		"FROM TEST.USERS WHERE EMAIL = '"+user.Email+"' AND PASSWORD = '"+user.Password+"'")
	if err != nil {
		log.Fatal("Error en la consulta")
		return err
	}

	for rows.Next() {
		var id int
		var username string
		err := rows.Scan(&id, &username)
		if err != nil {
			return err
		}
		user.Id = id
		user.Username = username
	}

	cookie := fiber.Cookie{
		Name: "user",
		Value: strconv.Itoa(user.Id),
		Expires: time.Now().Add(time.Hour*24), // 1 Día
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(user)
}

func User (c *fiber.Ctx) error {
	cookie := c.Cookies("user")
	id, _ := strconv.Atoi(cookie)

	rows, err := database.DB.Query("SELECT USERNAME, PASSWORD, FIRST_NAME, LAST_NAME, " +
		"DATE_BIRTH, DATE_REGISTER, EMAIL, PATH_PHOTO, FK_IDROLE " +
		"FROM TEST.USERS WHERE ID_USER = '"+cookie+"' ")
	if err != nil {
		fmt.Println("Error en la consulta")
		log.Fatal(err)
		return err
	}

	if id == 0 {
		return c.JSON(fiber.Map{
			"msg": "unauthenticated",
		})
	}

	var user models.User
	var username string
	var password string
	var firstName string
	var lastName string
	var dateBirth string
	var dateRegister string
	var email string
	var pathPhoto string
	var idRole int
	for rows.Next() {
		err := rows.Scan(&username, &password, &firstName, &lastName, &dateBirth, &dateRegister,
			&email, &pathPhoto, &idRole)
		if err != nil {
			return err
		}
	}
	user.Id = id
	user.Username = username
	user.Password = password
	user.First = firstName
	user.Last = lastName
	user.DateBirth = dateBirth
	user.DateRegister = dateRegister
	user.Email = email
	user.PathPhoto = pathPhoto
	user.IdRol = idRole

	return c.JSON(user)
}

func Logout (c *fiber.Ctx) error {
	cookie := fiber.Cookie{ // Eliminamos la cookie
		Name: "user",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"msg": "success",
	})
}