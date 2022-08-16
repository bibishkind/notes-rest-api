package handler

import "github.com/bibishkin/notes-rest-api/pkg/service"

type Handler struct {
	service service.Implementation
}

func NewHandler(service service.Implementation) *Handler {
	return &Handler{
		service: service,
	}
}
