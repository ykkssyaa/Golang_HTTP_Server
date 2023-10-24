package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	gateways := gateway.NewGateway(db, logger)

	logger.Info.Print("Creating Services.")
	services := service.NewService(gateways, logger)

	logger.Info.Print("Creating server.")

	port := viper.GetString("PORT")
	srv := server.NewHttpServer(services, logger, ":"+port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Err.Fatalf("error occured while running http server: \"%s\" \n", err.Error())
		}
	}()

	logger.Info.Print("Starting the server on port: " + port + "\n\n")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info.Println("Server Shutting Down.")

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Err.Fatalf("error occured while server shutting down: \"%s\" \n", err.Error())
	}

	logger.Info.Println("DB connection closing.")

	if err := db.Close(); err != nil {
		logger.Err.Fatalf("error occured on db connection close: \"%s\" \n", err.Error())
	}
}
