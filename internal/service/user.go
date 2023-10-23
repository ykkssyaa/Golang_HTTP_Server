package service

import (
	"errors"
	"testTask/internal/gateway"
	"testTask/internal/model"
	logger2 "testTask/pkg/logger"
)

type UserService interface {
	GetUsers(limit, offset int, filter model.UserFilter) ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	DeleteUser(id int) error
	UpdateUser(user model.User) error
}

type UserServiceImpl struct {
	repo   gateway.PostgresUserGateway
	api    gateway.UserThirdPartyApi
	logger *logger2.Logger
}

func (u UserServiceImpl) GetUsers(limit, offset int, filter model.UserFilter) ([]model.User, error) {
	return u.repo.GetUsers(limit, offset, filter)
}

func (u UserServiceImpl) CreateUser(user model.User) (model.User, error) {

	countries, err := u.api.GetCountry(user.Name)
	if err != nil {
		return model.User{}, err
	}
	user.Country = countryWithMaxProbability(countries)

	age, err := u.api.GetAge(user.Name)
	if err != nil {
		return model.User{}, err
	}
	user.Age = age

	gender, err := u.api.GetGender(user.Name)
	if err != nil {
		return model.User{}, err
	}
	user.Gender = model.Gender(gender)

	id, err := u.repo.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}
	user.Id = id

	return user, nil
}

func countryWithMaxProbability(countries []gateway.CountryProbability) string {

	if len(countries) == 0 {
		return ""
	}

	maxP := countries[0].Probability
	index := 0

	for i := 0; i < len(countries); i++ {
		if countries[i].Probability > maxP {
			index = i
		}
	}

	return countries[index].Country
}

func (u UserServiceImpl) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}

func (u UserServiceImpl) UpdateUser(user model.User) error {

	if user.Age < 0 {
		return errors.New("invalid age")
	}

	return u.repo.UpdateUser(user)
}
