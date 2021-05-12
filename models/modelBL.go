package models

type ModelBL struct {
	Firstname string `json:"nombre"`
	Lastname string `json:"apellido"`
	Password string `json:"password"`
	Username string `json:"username"`
	FootballPool [] struct{
		NameSeason string `json:"temporada"`
		Tier string `json:"tier"`
		Workdays [] struct{
			NameWorkday string `json:"jornada"`
			Events [] struct{
				NameSport string `json:"deporte"`
				DateEvent string `json:"fecha"`
				Visitor string `json:"visitante"`
				Local string `json:"local"`
				Prediction struct{
					UserVisitor int `json:"visitante"`
					UserLocal int `json:"local"`
				} `json:"prediccion"`
				Result  struct{
					RealVisitor int `json:"visitante"`
					RealLocal int `json:"local"`
				} `json:"resultado"`
			} `json:"predicciones"`
		} `json:"jornadas"`
	} `json:"resultados"`
}
