package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/models"
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
