package models

type User struct {
	Id        int
	Nick      string
	Email     string
	Firstname string
	LastName  string
}

type Result struct {
	Message string
}
