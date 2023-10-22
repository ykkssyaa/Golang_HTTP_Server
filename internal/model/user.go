package model

type gender string

const male gender = "male"
const female gender = "female"

type User struct {
	Id         int    `json:"-" db:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     gender `json:"gender"`
	Country    string `json:"country"`
}
