package handler

import "net/http"

func (h *HttpServer) CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("CreateUser"))
}

func (h *HttpServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}

func (h *HttpServer) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func (h *HttpServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUsers"))
}
