package models

type User struct {
	Id        int
	Username string
	Email string
	Password  string
	First     string
	Last      string
	Tier      string
	DateBirth string
	DateRegister string
	PathPhoto string
	IdRol int
	IdMembership int
}

type Role struct {
	IdRole int
	roleName string
}
