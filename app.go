package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	config "testTask/config"
	"testTask/internal/gateway"
	"testTask/internal/server"
	"testTask/internal/service"
	logger2 "testTask/pkg/logger"
)

func main() {

	logger := logger2.InitLogger()

	logger.Info.Print("Executing InitConfig.")

	if err := config.InitConfig(); err != nil {
		logger.Err.Fatalf(err.Error())
	}

	logger.Info.Print("Connecting to Postgres.")
	db, err := gateway.NewPostgresDB(viper.GetString("POSTGRES_DSN"))

	if err != nil {
		logger.Err.Fatalf(err.Error())
	}

	logger.Info.Print("Creating Gateways.")
	gateways := gateway.NewGateway(db)

	logger.Info.Print("Creating Services.")
	services := service.NewService(gateways)

	logger.Info.Print("Creating server.")

	port := viper.GetString("PORT")

	srv := server.NewHttpServer(services, logger, ":"+port)

	logger.Info.Print("Starting the server on port: " + port)

	logger.Err.Fatalf(srv.ListenAndServe().Error())

}
