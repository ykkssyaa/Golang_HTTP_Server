package model

type Gender string

const Male Gender = "male"
const female Gender = "female"

type User struct {
	Id         int    `json:"-" db:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     Gender `json:"Gender"`
	Country    string `json:"country"`
}
