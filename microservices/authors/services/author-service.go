package services

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/authors/models"
	"github.com/synapsis-library-management-server/microservices/authors/models/dto"
	"github.com/synapsis-library-management-server/microservices/authors/utils/pagination"
)

func (s *Service) CreateAuthor(req dto.CreateAuthorRequest) (string, error) {
	message := "Failed"

	author := req.ToModel()
	err := s.Repository.CreateAuthor(&author)
	if err != nil {
		log.Error().Err(err).Msg("[CreateAuthor] Service error creating author")
		return message, err
	}

	message = "Success"
	return message, nil
}

func (s *Service) GetAuthorsByFilter(req dto.GetAuthorsByFilterRequest) ([]dto.GetAuthorsByFilterResponse, dto.PaginationResponse, error) {
	authors, totalAuthors, err := s.Repository.GetAuthorsByFilter(models.Filter{
		Pagination: models.Pagination{
			Page:     int(req.Page),
			PageSize: int(req.PageSize),
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetAuthorsByFilter] Service error getting authors")
		return []dto.GetAuthorsByFilterResponse{}, dto.PaginationResponse{}, err
	}

	responses := dto.BuildGetAuthorsByFilterResponse(authors)
	metadata := pagination.CalculatePaginationMetadata(req.Page, req.PageSize, totalAuthors)
	return responses, metadata, nil
}
