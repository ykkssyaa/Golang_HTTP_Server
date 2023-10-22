package gateway

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testTask/internal/model"
)

type PostgresUserGateway interface {
	GetUsers(limit, offset int, filter model.UserFilter) ([]model.User, error)
	CreateUser(user model.User) (int, error)
	DeleteUser(id int) error
	UpdateUser(user model.User) error
}

type PostgresUserGatewayImpl struct {
	db *sqlx.DB
}

func (u PostgresUserGatewayImpl) GetUsers(limit, offset int, filter model.UserFilter) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u PostgresUserGatewayImpl) CreateUser(user model.User) (int, error) {

	var id int

	createQuery := fmt.Sprintf("INSERT INTO %s "+
		"(name, surname, patronymic, age, country, gender) "+
		"VALUES ($1, $2, $3, $4, $5, $6)"+
		"RETURNING id", usersTable)

	tx, err := u.db.DB.Begin()

	if err != nil {
		return 0, err
	}
	row := tx.QueryRow(createQuery, user.Name, user.Surname, user.Patronymic, user.Age, user.Country, user.Gender)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, err
}

func (u PostgresUserGatewayImpl) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (u PostgresUserGatewayImpl) UpdateUser(user model.User) error {
	//TODO implement me
	panic("implement me")
}
