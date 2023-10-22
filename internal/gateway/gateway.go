package gateway

import "github.com/jmoiron/sqlx"

type Gateways struct {
	UserGateway       PostgresUserGateway
	UserThirdPartyApi UserThirdPartyApi
}

func NewGateway(db *sqlx.DB) *Gateways {
	return &Gateways{
		UserGateway:       PostgresUserGatewayImpl{db: db},
		UserThirdPartyApi: UserThirdPartyApiImpl{},
	}
}
