package data

type Person struct {
	Id       string
	Name   string
	Surname string
	Team   string
	Age     int
	Word  string
}

//Enum type for teams
const (
		RED string = "rojo"
		BLUE string = "azul"
	)