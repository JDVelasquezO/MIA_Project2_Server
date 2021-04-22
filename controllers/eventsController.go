package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/models"
)

func GetEvents(c *fiber.Ctx) error {

	query := "SELECT ID_EVENT, DATE_OF_GAME, COLOR, IDTEAM, NAME_TEAM, NAME_CLASSIFICATION, NAME_SPORT FROM EVENT " +
		"INNER JOIN STATUS_EVENT SE on SE.IDSTATUSEVENT = EVENT.FK_IDSTATUSEVENT " +
		"INNER JOIN EVENT_HAS_TEAM EHT on EVENT.ID_EVENT = EHT.FK_IDEVENT " +
		"INNER JOIN TEAM T on EHT.FK_IDTEAM = T.IDTEAM " +
		"INNER JOIN CLASSIFICATION C2 on C2.ID_CLASSIFICATION = EHT.FK_IDCLASSIFICATION " +
		"INNER JOIN SPORT S on T.FK_IDSPORT = S.ID_SPORT"

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

	for rows.Next() {
		var event models.Event
		var team models.Team

		err := rows.Scan(&idEvent, &dateGame, &color, &idTeam, &nameTeam, &nameClass, &nameSport)
		if err != nil {
			return err
		}

		event.IdEvent = idEvent
		event.Color = color
		event.DateOfGame = dateGame
			team.IdTeam = idTeam
			team.NameTeam = nameTeam
			team.Classification = nameClass
		event.NameSport = nameSport

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
