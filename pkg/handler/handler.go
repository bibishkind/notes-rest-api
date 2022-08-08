package handler

import (
	"github.com/bibishkin/bi-notes-rest-api/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/auth/sign-up", h.signUp).Methods("POST")
	r.HandleFunc("/auth/sign-in", h.signIn).Methods("POST")

	return r
}
