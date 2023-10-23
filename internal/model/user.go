package model

type Gender string

type User struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     Gender `json:"Gender"`
	Country    string `json:"country"`
}
