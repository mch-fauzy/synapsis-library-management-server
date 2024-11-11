package handlers

import (
	"github.com/synapsis-library-management-server/microservices/categories/services"
)

type Handler struct {
	Service *services.Service
}

// NewHandler is the constructor for Handler
func NewHandler(service *services.Service) Handler {
	return Handler{
		Service: service,
	}
}
