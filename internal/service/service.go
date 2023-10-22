package service

import "testTask/internal/gateway"

type Services struct {
	UserService UserService
}

func NewService(gateways *gateway.Gateways) *Services {
	return &Services{
		UserService: UserServiceImpl{
			repo: gateways.UserGateway,
			api:  gateways.UserThirdPartyApi,
		},
	}
}
