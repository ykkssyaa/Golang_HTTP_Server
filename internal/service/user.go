package service

import (
	"testTask/internal/gateway"
	"testTask/internal/model"
)

type UserService interface {
	GetUsers(limit, offset int, filter model.UserFilter) ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	DeleteUser(id int) error
	UpdateUser(user model.User) error
}

type UserServiceImpl struct {
	repo gateway.PostgresUserGateway
	api  gateway.UserThirdPartyApi
}

func (u UserServiceImpl) GetUsers(limit, offset int, filter model.UserFilter) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) CreateUser(user model.User) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) UpdateUser(user model.User) error {
	//TODO implement me
	panic("implement me")
}
