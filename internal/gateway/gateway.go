package gateway

import (
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

type Gateways struct {
	UserGateway       PostgresUserGateway
	UserThirdPartyApi UserThirdPartyApi
}

func NewGateway(db *sqlx.DB) *Gateways {
	return &Gateways{
		UserGateway:       PostgresUserGatewayImpl{db: db},
		UserThirdPartyApi: UserThirdPartyApiImpl{client: &http.Client{Timeout: 30 * time.Second}},
	}
}
