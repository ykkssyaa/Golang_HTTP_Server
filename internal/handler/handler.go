package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"testTask/internal/service"
	"time"
)

type HttpServer struct {
	services *service.Services
}

func NewHttpServer(services *service.Services, addr string) *http.Server {

	server := &HttpServer{services: services}

	r := mux.NewRouter()
	r.HandleFunc("/api/user", server.GetUsers).Methods("GET")
	r.HandleFunc("/api/user", server.CreateUser).Methods("POST")
	r.HandleFunc("/api/user", server.DeleteUser).Methods("DELETE")
	r.HandleFunc("api/user", server.UpdateUser).Methods("UPDATE")

	return &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Addr:           addr,
	}
}
