package models

type Event struct {
	IdEvent int
	DateOfGame string
	Color string
	Teams []Team
	NameSport string
}
