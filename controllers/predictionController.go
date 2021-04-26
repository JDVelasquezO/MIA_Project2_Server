package controllers

/*func PostPrediction (c *fiber.Ctx) error {
	var data map[string]int // key: value
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	userRes2 := data["userRes1"]
	userRes3 := data["userRes3"]
	idMembership := data["id_membership"]
	idEvent := data["id_event"]

	// OBTENER IDS DE LAS PREDICCIONES
	queryGetIdsPrediction := "SELECT FK_IDPREDICTION FROM EVENT_HAS_TEAM " +
		"WHERE FK_IDEVENT = "+strconv.Itoa(idEvent)+" "
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
		"USER_RESULT = "++" " +
		"WHERE ID_PREDICTION = 7;"

	return c.JSON(fiber.Map{
		"res1": idsPredictions,
	})
}*/
