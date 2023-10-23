package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"testTask/internal/model"
)

const (
	defaultLimit  = 10
	defaultOffset = 0
)

func (h *HttpServer) CreateUser(w http.ResponseWriter, r *http.Request) {

	h.logger.Info.Println("Invoked CreateUser of Server")

	if r.Header.Get("Content-Type") != "application/json" {
		h.logger.Err.Println("Error: Content Type is not application/json")
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var user model.User
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	h.logger.Info.Println("Decoding json from request body.")
	err := decoder.Decode(&user)

	if err != nil {
		h.logger.Err.Println(err.Error())
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	h.logger.Info.Println("Request body: ", user)

	if len(user.Name) == 0 || len(user.Surname) == 0 {
		h.logger.Err.Println("Name or Surname is empty")
		errorResponse(w, "Name or Surname is empty", http.StatusBadRequest)
		return
	}

	h.logger.Info.Println("Invoking UserService.CreateUser")

	newUser, err := h.services.UserService.CreateUser(user)

	if err != nil {
		h.logger.Err.Println(err.Error())
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.logger.Info.Println("Result user: ", newUser)
	h.logger.Info.Println("Sending result.\n\n")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (h *HttpServer) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpServer) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpServer) GetUsers(w http.ResponseWriter, r *http.Request) {

	h.logger.Info.Println("Invoked GetUsers of Server")

	var limit, offset int

	// Getting limit from URL params
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil {
			h.logger.Err.Println("Invalid limit parameter: ", limitStr)
			errorResponse(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}

		limit = l
	} else {
		limit = defaultLimit
	}

	// Getting offset from URL params
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		off, err := strconv.Atoi(offsetStr)
		if err != nil {
			h.logger.Err.Println("Invalid offset parameter: ", offsetStr)
			errorResponse(w, "Invalid offset parameter", http.StatusBadRequest)
			return
		}
		offset = off
	} else {
		offset = defaultOffset
	}

	filter, err := CreateUserFilter(r)
	if err != nil {
		errorResponse(w, "Invalid age parameter", http.StatusBadRequest)
		return
	}

	h.logger.Info.Println("Invoking UserService.GetUsers")

	users, err := h.services.UserService.GetUsers(limit, offset, filter)

	if err != nil {
		h.logger.Err.Println(err.Error())
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)

}

func CreateUserFilter(r *http.Request) (model.UserFilter, error) {
	var filter model.UserFilter

	// Getting name from URL params
	name := r.URL.Query().Get("name")
	filter.Name = name

	// Getting surname from URL params
	surname := r.URL.Query().Get("surname")
	filter.Surname = surname

	// Getting patronymic from URL params
	patronymic := r.URL.Query().Get("patronymic")
	filter.Patronymic = patronymic

	// Getting age from URL params
	ageStr := r.URL.Query().Get("age")

	if ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			return model.UserFilter{}, err
		}
		filter.Age = age
	}

	// Getting country from URL params
	country := r.URL.Query().Get("country")
	filter.Country = country

	// Getting gender from URL params
	gender := r.URL.Query().Get("gender")
	filter.Gender = model.Gender(gender)

	return filter, nil
}
