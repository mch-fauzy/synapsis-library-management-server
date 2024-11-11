package services

import (
	"github.com/synapsis-library-management-server/microservices/authors/repositories"
)

type Service struct {
	Repository *repositories.Repository
}

// NewService is the constructor for Service
func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}
