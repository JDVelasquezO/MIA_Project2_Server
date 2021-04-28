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

func Register(c *fiber.Ctx) error {
	var data map[string]string // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	IdRole, _ := strconv.Atoi(data["idRole"])
	user := models.User{
		Username:     data["username"],
		Email:        data["email"],
		Password:     data["pass"],
		First:        data["first"],
		Last:         data["last"],
		DateBirth:    data["birth"],
		DateRegister: data["register"],
		PathPhoto:    data["photo"],
		IdRol:        IdRole,
	}

	if user.PathPhoto == "" {
		user.PathPhoto = "https://i.pinimg.com/originals/e5/0e/30/e50e3015eb9d4fb800bac9c53815e1f6.png"
	}

	query := "INSERT INTO TEST.USERS (USERNAME, PASSWORD, FIRST_NAME, LAST_NAME, DATE_BIRTH," +
		"DATE_REGISTER, EMAIL, PATH_PHOTO, FK_IDROLE) VALUES ('" + user.Username +
		"', '" + user.Password + "', '" + user.First + "', '" + user.Last + "', TO_DATE('" + user.DateBirth + "', 'yyyy/mm/dd')," +
		"TO_DATE('" + user.DateRegister + "', 'yyyy/mm/dd'), '" + user.Email + "', '" + user.PathPhoto + "', 4)"

	_, err := database.DB.Query(query)
	// e = fmt.Errorf("Mal formato para email: %v", user.Email)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	var user models.User
	user.Username = data["username"]
	user.Password = data["pass"]
	rows, err := database.DB.Query("SELECT ID_USER " +
		"FROM TEST.USERS WHERE USERNAME = '" + user.Username + "' AND PASSWORD = '" + user.Password + "'")
	if err != nil {
		log.Fatal("Error en la consulta")
		return err
	}

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
		user.Id = id
	}

	if user.Id != 0 {
		cookie := fiber.Cookie{
			Name:     "user",
			Value:    strconv.Itoa(user.Id),
			Expires:  time.Now().Add(time.Hour * 24), // 1 Día
			HTTPOnly: true,
		}
		c.Cookie(&cookie)
	}

	return c.JSON(user)
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("user")
	id, _ := strconv.Atoi(cookie)

	rows, err := database.DB.Query("SELECT USERNAME, PASSWORD, FIRST_NAME, LAST_NAME, " +
		"DATE_BIRTH, DATE_REGISTER, EMAIL, PATH_PHOTO, FK_IDROLE " +
		"FROM TEST.USERS WHERE ID_USER = '" + cookie + "' ")
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

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{ // Eliminamos la cookie
		Name:     "user",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"msg": "success",
	})
}
