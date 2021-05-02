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

	query := "SELECT ID_EVENT, DATE_OF_GAME, COLOR, IDTEAM, NAME_CLASSIFICATION, " +
		"NAME_TEAM, REAL_RESULT, NAME_SPORT, C3.NAME_COLOR " +
		"FROM EVENT " +
		"INNER JOIN STATUS_EVENT SE on SE.IDSTATUSEVENT = EVENT.FK_IDSTATUSEVENT " +
		"INNER JOIN EVENT_HAS_TEAM EHT on EVENT.ID_EVENT = EHT.FK_IDEVENT " +
		"INNER JOIN TEAM T on EHT.FK_IDTEAM = T.IDTEAM " +
		"INNER JOIN CLASSIFICATION C2 on C2.ID_CLASSIFICATION = EHT.FK_IDCLASSIFICATION " +
		"INNER JOIN SPORT S on T.FK_IDSPORT = S.ID_SPORT " +
		"INNER JOIN COLOR C3 on S.FK_IDCOLOR = C3.ID_COLOR " +
		"GROUP BY ID_EVENT, DATE_OF_GAME, COLOR, IDTEAM, NAME_CLASSIFICATION, " +
		"NAME_TEAM, REAL_RESULT, NAME_SPORT, C3.NAME_COLOR " +
		"ORDER BY ID_EVENT ASC"

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
	var idTeam int
	var nameTeam string
	var nameClass string
	var nameSport string
	var colorSport string
	var realRes int

	for rows.Next() {
		var event models.Event
		var team models.Team

		err := rows.Scan(&idEvent, &dateGame, &color, &idTeam, &nameClass, &nameTeam,
			&realRes, &nameSport, &colorSport)
		if err != nil {
			return err
		}

		event.IdEvent = idEvent
		event.Color = color
		event.DateOfGame = dateGame
		team.IdTeam = idTeam
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

	query := "SELECT ID_EVENT, DATE_OF_GAME, COLOR, IDTEAM, NAME_CLASSIFICATION, " +
		"NAME_TEAM, USER_RESULT, REAL_RESULT, NAME_SPORT, C3.NAME_COLOR " +
		"FROM EVENT " +
		"INNER JOIN STATUS_EVENT SE on SE.IDSTATUSEVENT = EVENT.FK_IDSTATUSEVENT " +
		"INNER JOIN EVENT_HAS_TEAM EHT on EVENT.ID_EVENT = EHT.FK_IDEVENT " +
		"INNER JOIN TEAM T on EHT.FK_IDTEAM = T.IDTEAM " +
		"INNER JOIN CLASSIFICATION C2 on C2.ID_CLASSIFICATION = EHT.FK_IDCLASSIFICATION " +
		"INNER JOIN SPORT S on T.FK_IDSPORT = S.ID_SPORT " +
		"INNER JOIN PREDICTION P on EHT.FK_IDPREDICTION = P.ID_PREDICTION " +
		"INNER JOIN COLOR C3 on S.FK_IDCOLOR = C3.ID_COLOR " +
		"WHERE ID_EVENT = "+strconv.Itoa(paramIdEvent)+" " +
		"AND FK_IDMEMBERSHIP = "+strconv.Itoa(idMembership)+" "

	var event models.Event
	event = executeQuery(query)
	if event.IdEvent == 0 {
		// println("Entro aqui")
		newQuery := "SELECT IDTEAM, ID_EVENT, DATE_OF_GAME, COLOR, NAME_CLASSIFICATION, " +
			"NAME_TEAM, NAME_SPORT, C3.NAME_COLOR " +
			"FROM EVENT " +
			"INNER JOIN STATUS_EVENT SE on SE.IDSTATUSEVENT = EVENT.FK_IDSTATUSEVENT " +
			"INNER JOIN EVENT_HAS_TEAM EHT on EVENT.ID_EVENT = EHT.FK_IDEVENT " +
			"INNER JOIN TEAM T on EHT.FK_IDTEAM = T.IDTEAM " +
			"INNER JOIN CLASSIFICATION C2 on C2.ID_CLASSIFICATION = EHT.FK_IDCLASSIFICATION " +
			"INNER JOIN SPORT S on T.FK_IDSPORT = S.ID_SPORT " +
			"INNER JOIN COLOR C3 on S.FK_IDCOLOR = C3.ID_COLOR " +
			"WHERE ID_EVENT = "+strconv.Itoa(paramIdEvent)+" " +
			"GROUP BY IDTEAM, ID_EVENT, DATE_OF_GAME, COLOR, NAME_CLASSIFICATION, " +
			"NAME_TEAM, REAL_RESULT, NAME_SPORT, C3.NAME_COLOR"
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
	var idTeam int
	var nameTeam string
	var nameClass string
	var nameSport string
	var colorSport string
	var userRes int
	var realRes int

	var event models.Event
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&idEvent, &dateGame, &color, &idTeam, &nameClass, &nameTeam,
			&userRes, &realRes, &nameSport, &colorSport)
		if err != nil {
			println(err)
		}

		event.IdEvent = idEvent
		event.Color = color
		event.DateOfGame = dateGame
		team.IdTeam = idTeam
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
	var idTeam int
	var nameTeam string
	var nameClass string
	var nameSport string
	var colorSport string

	var event models.Event
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&idTeam, &idEvent, &dateGame, &color, &nameClass, &nameTeam,
			&nameSport, &colorSport)
		if err != nil {
			println(err)
		}

		event.IdEvent = idEvent
		event.Color = color
		event.DateOfGame = dateGame
		team.IdTeam = idTeam
		team.NameTeam = nameTeam
		team.Classification = nameClass
		event.NameSport = nameSport
		event.ColorSport = colorSport
		event.Teams = append(event.Teams, team)
	}
	event.Teams[0].UserResult = -1
	event.Teams[1].UserResult = -1
	return event
}