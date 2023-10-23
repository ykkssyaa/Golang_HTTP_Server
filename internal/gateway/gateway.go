package gateway

import (
	"github.com/jmoiron/sqlx"
	"net/http"
	logger2 "testTask/pkg/logger"
	"time"
)

type Gateways struct {
	UserGateway       PostgresUserGateway
	UserThirdPartyApi UserThirdPartyApi
}

func NewGateway(db *sqlx.DB, logger *logger2.Logger) *Gateways {
	return &Gateways{
		UserGateway:       PostgresUserGatewayImpl{db: db, logger: logger},
		UserThirdPartyApi: UserThirdPartyApiImpl{client: &http.Client{Timeout: 30 * time.Second}, logger: logger},
	}
}
