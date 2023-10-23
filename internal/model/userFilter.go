package model

type UserFilter struct {
	Id         int
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     Gender
	Country    string
}
