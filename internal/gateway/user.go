package gateway

import (
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
	//TODO implement me
	panic("implement me")
}

func (u PostgresUserGatewayImpl) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (u PostgresUserGatewayImpl) UpdateUser(user model.User) error {
	//TODO implement me
	panic("implement me")
}
