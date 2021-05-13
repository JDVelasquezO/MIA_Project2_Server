package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/models"
	"strconv"
)

func GetEvents(c *fiber.Ctx) error {
	cookie := c.Cookies("user")

	queryGetMembership := "SELECT ID_MEMBERSHIP, FK_IDTIER FROM MEMBERSHIP " +
		"WHERE FK_IDUSER = "+cookie+" "

	var idMembership int
	var fkIdTier int
	rows0, _ := database.DB.Query(queryGetMembership)
	for rows0.Next() {
		err0 := rows0.Scan(&idMembership, &fkIdTier)
		if err0 != nil {
			return err0
		}
	}

	if fkIdTier == 4 {
		return nil
	}

	query := "SELECT ID_EVENT, DATE_OF_GAME, COLOR, PLAYER, NAME_CLASSIFICATION, " +
		"NAME_SPORT, C3.COD_HEX, REAL_RESULT " +
		"FROM EVENT " +
		"INNER JOIN STATUS_EVENT SE on SE.IDSTATUSEVENT = EVENT.FK_IDSTATUSEVENT " +
		"INNER JOIN CLASSIFICATION C2 on C2.ID_CLASSIFICATION = EVENT.FK_IDCLASS " +
		"INNER JOIN SPORT S2 on S2.ID_SPORT = EVENT.FK_IDSPORT " +
		"INNER JOIN COLOR C3 on S2.FK_IDCOLOR = C3.ID_COLOR " +
		"ORDER BY ID_EVENT"

	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Error en la consulta")
		log.Fatal(err)
		return err
	}

	var events []models.Event
	var idEvent int
	var dateGame string
	var color string
	var nameTeam string
	var nameClass string
	var nameSport string
	var colorSport string
	var realRes int

	for rows.Next() {
		var event models.Event
		var team models.Team

		err := rows.Scan(&idEvent, &dateGame, &color, &nameTeam, &nameClass,
			&nameSport, &colorSport, &realRes)
		if err != nil {
			return err
		}

		event.IdEvent = idEvent
		event.Color = color
		event.DateOfGame = dateGame
		team.NameTeam = nameTeam
		team.Classification = nameClass
		team.RealResult = realRes
		event.NameSport = nameSport
		event.ColorSport = colorSport

		event.Teams = append(event.Teams, team)
		events = append(events, event)
	}

	var teams []models.Team
	for i, e := range events {
		if i % 2 != 0 {
			teams = append(teams, e.Teams[0])
			// println(i , e.Teams[0].NameTeam)
		}
	}

	var newEvents []models.Event
	for i, e := range events {
		if i % 2 == 0 {
			e.Teams = append(e.Teams, teams[e.IdEvent-1])
			newEvents = append(newEvents, e)
		}
	}

	return c.JSON(newEvents)
}

func GetEvent(c *fiber.Ctx) error {
	cookie := c.Cookies("user")
	paramIdEvent, _ := c.ParamsInt("id")

	queryGetMembership := "SELECT ID_MEMBERSHIP FROM MEMBERSHIP " +
		"WHERE FK_IDUSER = "+cookie+" "

	var idMembership int
	rows, err := database.DB.Query(queryGetMembership)
	if err != nil {
		return err
	}

	for rows.Next() {
		err := rows.Scan(&idMembership)
		if err != nil {
			return err
		}
	}

	query := "SELECT ID_EVENT, DATE_OF_GAME, COLOR, NAME_CLASSIFICATION, PLAYER, " +
		"USER_RESULT, REAL_RESULT, NAME_SPORT, C3.NAME_COLOR " +
		"FROM EVENT " +
		"INNER JOIN STATUS_EVENT SE on SE.IDSTATUSEVENT = EVENT.FK_IDSTATUSEVENT " +
		"INNER JOIN PREDICTION P on P.FK_IDEVENT = EVENT.ID_EVENT " +
		"INNER JOIN SPORT S2 on S2.ID_SPORT = EVENT.FK_IDSPORT " +
		"INNER JOIN COLOR C3 on C3.ID_COLOR = S2.FK_IDCOLOR " +
		"INNER JOIN CLASSIFICATION C2 on C2.ID_CLASSIFICATION = P.FK_IDCLASS AND EVENT.FK_IDCLASS = C2.ID_CLASSIFICATION " +
		"WHERE FK_IDEVENT = "+strconv.Itoa(paramIdEvent)+" " +
		"AND FK_IDMEMBERSHIP = "+strconv.Itoa(idMembership)+" " +
		"GROUP BY P.FK_IDCLASS, ID_EVENT, DATE_OF_GAME, COLOR, NAME_CLASSIFICATION, " +
		"PLAYER, USER_RESULT, REAL_RESULT, NAME_SPORT, C3.NAME_COLOR"

	var event models.Event
	event = executeQuery(query)
	if event.IdEvent == 0 {
		// println("Entro aqui")
		newQuery := "SELECT ID_EVENT, DATE_OF_GAME, COLOR, PLAYER, NAME_SPORT, " +
			"FK_IDCLASS, NAME_CLASSIFICATION, C3.NAME_COLOR " +
			"FROM EVENT " +
			"INNER JOIN STATUS_EVENT SE on SE.IDSTATUSEVENT = EVENT.FK_IDSTATUSEVENT " +
			"INNER JOIN SPORT S2 on S2.ID_SPORT = EVENT.FK_IDSPORT " +
			"INNER JOIN CLASSIFICATION C2 on C2.ID_CLASSIFICATION = EVENT.FK_IDCLASS " +
			"INNER JOIN COLOR C3 on C3.ID_COLOR = S2.FK_IDCOLOR " +
			"WHERE ID_EVENT = "+strconv.Itoa(paramIdEvent)+" " +
			"GROUP BY ID_EVENT, DATE_OF_GAME, COLOR, " +
			"PLAYER, NAME_SPORT, FK_IDCLASS, NAME_CLASSIFICATION, C3.NAME_COLOR"
		event = executeQuery2(newQuery)
	}

	return c.JSON(event)
}

func executeQuery (query string) models.Event {
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Error en la consulta")
		log.Fatal(err)
	}

	var idEvent int
	var dateGame string
	var color string
	var nameTeam string
	var nameClass string
	var nameSport string
	var colorSport string
	var userRes int
	var realRes int

	var event models.Event
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&idEvent, &dateGame, &colorSport, &nameClass, &nameTeam,
			&userRes, &realRes, &nameSport, &color)
		if err != nil {
			println(err)
		}

		event.IdEvent = idEvent
		event.Color = color
		event.DateOfGame = dateGame
			team.NameTeam = nameTeam
			team.Classification = nameClass
			team.RealResult = realRes
			team.UserResult = userRes
		event.NameSport = nameSport
		event.ColorSport = colorSport
		event.Teams = append(event.Teams, team)
	}
	return event
}

func executeQuery2 (query string) models.Event {
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Error en la consulta")
		log.Fatal(err)
	}

	var idEvent int
	var dateGame string
	var color string
	var nameTeam string
	var idClass int
	var nameClass string
	var nameSport string
	var colorSport string

	var event models.Event
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&idEvent, &dateGame, &color, &nameTeam, &nameSport,
			&idClass, &nameClass, &colorSport)
		if err != nil {
			println("Error", err.Error())
		}

		event.IdEvent = idEvent
		event.Color = color
		event.DateOfGame = dateGame
		team.NameTeam = nameTeam
		team.Classification = nameClass
		team.IdClass = idClass
		event.NameSport = nameSport
		event.ColorSport = colorSport
		event.Teams = append(event.Teams, team)
	}
	event.Teams[0].UserResult = -1
	event.Teams[1].UserResult = -1
	return event
}

func PostEvent(c *fiber.Ctx) error {
	// Recuperar conteo de eventos
	query := "SELECT COUNT(*) FROM EVENT"
	rows, _ := database.DB.Query(query)
	var idEvent int
	for rows.Next() {
		err := rows.Scan(&idEvent)
		if err != nil {
			println(err)
			return err
		}
	}

	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	// var event models.Event
	newIdEvent := (idEvent + 2)/2
	eventDate := data["dateEvent"]
	fkIdStatus := data["idStatus"]
	fkIdWd := data["idWd"]
	fkIDSport := data["idSport"]
	fkIdClass1 := data["idClass1"]
	fkIdClass2 := data["idClass2"]
	player1 := data["player1"]
	player2 := data["player2"]

	// Insertar evento nuevo
	query2 := "INSERT ALL " +
		"INTO EVENT (ID_EVENT, DATE_OF_GAME, FK_IDSTATUSEVENT, FK_IDWORKINGDAY, FK_IDSPORT, FK_IDCLASS, REAL_RESULT, PLAYER) " +
		"VALUES ("+strconv.Itoa(newIdEvent)+", TO_DATE('"+eventDate+"', 'yyyy/mm/dd hh24:mi'), "+fkIdStatus+", " +
		" "+fkIdWd+", "+fkIDSport+", "+fkIdClass1+", 0, '"+player1+"' ) " +
		"INTO EVENT (ID_EVENT, DATE_OF_GAME, FK_IDSTATUSEVENT, FK_IDWORKINGDAY, FK_IDSPORT, FK_IDCLASS, REAL_RESULT, PLAYER) " +
		"VALUES ("+strconv.Itoa(newIdEvent)+", TO_DATE('"+eventDate+"', 'yyyy/mm/dd hh24:mi'), "+fkIdStatus+", " +
		" "+fkIdWd+", "+fkIDSport+", "+fkIdClass2+", 0, '"+player2+"' ) " +
		"SELECT 1 FROM DUAL"

	_, err := database.DB.Query(query2)
	if err != nil {
		println(err.Error())
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "success",
	})
}

func UpdateResults (c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	idEvent := data["idEvent"]
	idClass1 := data["idClass1"]
	idClass2 := data["idClass2"]
	res1 := data["res1"]
	res2 := data["res2"]
	println(idClass2, idClass1, idEvent, res1, res2)

	query := "UPDATE EVENT SET REAL_RESULT = "+res1+", FK_IDSTATUSEVENT = 2 " +
		"WHERE ID_EVENT = "+idEvent+" AND FK_IDCLASS = "+idClass1+" "
	query2 := "UPDATE EVENT SET REAL_RESULT = "+res2+", FK_IDSTATUSEVENT = 2 " +
		"WHERE ID_EVENT = "+idEvent+" AND FK_IDCLASS = "+idClass2+" "

	_, err := database.DB.Query(query)
	_, err2 := database.DB.Query(query2)
	if err != nil || err2 != nil {
		println(err.Error())
		println(err2.Error())
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "success",
	})
}