package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"testTask/internal/model"
)

func (h *HttpServer) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var user model.User
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&user)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	newUser, err := h.services.UserService.CreateUser(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (h *HttpServer) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpServer) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpServer) GetUsers(w http.ResponseWriter, r *http.Request) {

}
