package handler

import service2 "github.com/bibishkin/bi-notes-rest-api/pkg/service"

type Handler struct {
	service *service2.Service
}

func NewHandler(service *service2.Service) *Handler {
	return &Handler{
		service: service,
	}
}
