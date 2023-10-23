package service

import (
	"testTask/internal/gateway"
	logger2 "testTask/pkg/logger"
)

type Services struct {
	UserService UserService
}

func NewService(gateways *gateway.Gateways, logger *logger2.Logger) *Services {
	return &Services{
		UserService: UserServiceImpl{
			repo:   gateways.UserGateway,
			api:    gateways.UserThirdPartyApi,
			logger: logger,
		},
	}
}
