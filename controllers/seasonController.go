package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/models"
	"strconv"
)

func GetActualSeason (c *fiber.Ctx) error {
	query := "SELECT SEASON.NAME, COUNT(*) Cantidad_Participantes, ( " +
		"SELECT SUM(TIER_PRICE) " +
		"FROM MEMBERSHIP " +
		"INNER JOIN TIER T on T.IDTIER = MEMBERSHIP.FK_IDTIER " +
		") AS CAPITAL " +
		"FROM SEASON " +
		"INNER JOIN MEMBERSHIP M ON SEASON.ID_SEASON = M.FK_SEASON " +
		"INNER JOIN TIER T on T.IDTIER = M.FK_IDTIER " +
		"WHERE  ( " +
		"SELECT CURRENT_DATE FROM DUAL ) " +
		"BETWEEN SEASON.START_DATE AND SEASON.END_DATE " +
		"AND M.FK_IDTIER != 4 " +
		"GROUP BY SEASON.NAME, M.FK_IDTIER " +
		"ORDER BY M.FK_IDTIER DESC"
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal("Error en la consulta")
		return err
	}

	var season models.Season
	var nameSeason string
	var capital float32
	var quant int
	var quants []int
	for rows.Next() {
		err := rows.Scan(&nameSeason, &quant, &capital)
		if err != nil {
			return err
		}

		quants = append(quants, quant)
	}
	season.NameSeason = nameSeason
	season.Capital = capital
	season.QuantityBronze = quants[0]
	season.QuantitySilver = quants[1]
	season.QuantityGold = quants[2]

	return c.JSON(season)
}

func GetParticipants (c *fiber.Ctx) error {
	query := "SELECT ID_USER, USERNAME, ID_MEMBERSHIP FROM USERS " +
		"INNER JOIN MEMBERSHIP M on USERS.ID_USER = M.FK_IDUSER"

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal("Error en la consulta")
		return err
	}

	var user models.User
	var idUser int
	var username string
	var idMembership int
	var users []models.User
	for rows.Next() {
		err := rows.Scan(&idUser, &username, &idMembership)
		if err != nil {
			return err
		}
		user.Id = idUser
		user.Username = username
		user.IdMembership = idMembership

		users = append(users, user)
	}

	return c.JSON(users)
}

func GetEventsOfParticipant (c *fiber.Ctx) error {
	paramIdMembership, _ := c.ParamsInt("id")

	query := "SELECT ID_EVENT, NAME_SPORT, PLAYER, USER_RESULT, REAL_RESULT, DATE_OF_GAME, NAME_CLASSIFICATION " +
		"FROM EVENT " +
		"INNER JOIN PREDICTION ON PREDICTION.FK_IDEVENT = EVENT.ID_EVENT " +
		"AND PREDICTION.FK_IDCLASS = EVENT.FK_IDCLASS " +
		"INNER JOIN SPORT S2 on S2.ID_SPORT = EVENT.FK_IDSPORT " +
		"INNER JOIN CLASSIFICATION C2 on C2.ID_CLASSIFICATION = EVENT.FK_IDCLASS " +
		"WHERE FK_IDMEMBERSHIP = "+strconv.Itoa(paramIdMembership)+" " +
		"ORDER BY ID_EVENT, ID_CLASSIFICATION"

	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Error en la consulta")
		log.Fatal(err)
		return err
	}

	var idEvent int
	var dateGame string
	var nameTeam string
	var nameSport string
	var nameClass string
	var realRes int
	var userRes int

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var team models.Team

		err := rows.Scan(&idEvent, &nameSport, &nameTeam, &userRes, &realRes, &dateGame, &nameClass)
		if err != nil {
			println("Error en la consulta")
			return err
		}

		event.IdEvent = idEvent
		event.NameSport = nameSport
		team.NameTeam = nameTeam
		team.UserResult = userRes
		team.RealResult = realRes
		team.Classification = nameClass
		event.DateOfGame = dateGame

		event.Teams = append(event.Teams, team)
		events = append(events, event)
	}

	return c.JSON(events)
}