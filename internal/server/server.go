package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"testTask/internal/service"
	logger2 "testTask/pkg/logger"
	"time"
)

type HttpServer struct {
	services *service.Services
	logger   *logger2.Logger
}

func NewHttpServer(services *service.Services, logger *logger2.Logger, addr string) *http.Server {

	server := &HttpServer{services: services, logger: logger}

	r := mux.NewRouter()
	r.HandleFunc("/api/user", server.GetUsers).Methods("GET")
	r.HandleFunc("/api/user", server.CreateUser).Methods("POST")
	r.HandleFunc("/api/user", server.DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/user", server.UpdateUser).Methods("PATCH")

	return &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Addr:           addr,
	}
}
