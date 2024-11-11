package services

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/authors/models"
	"github.com/synapsis-library-management-server/microservices/authors/models/dto"
	"github.com/synapsis-library-management-server/microservices/authors/utils/failure"
	"github.com/synapsis-library-management-server/microservices/authors/utils/pagination"
	"gorm.io/gorm"
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

func (s *Service) GetAuthorById(req dto.GetAuthorByIdRequest) (dto.AuthorResponse, error) {

	author, err := s.Repository.GetAuthorById(models.AuthorPrimaryId{Id: req.Id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = failure.NotFound("Author not found")
			return dto.AuthorResponse{}, err
		}

		log.Error().Err(err).Msg("[GetAuthorById] Service error retrieving author by id")
		return dto.AuthorResponse{}, err
	}

	response := dto.NewAuthorResponse(author)
	return response, nil
}

func (s *Service) GetAuthorsByFilter(req dto.GetAuthorsByFilterRequest) ([]dto.AuthorResponse, dto.PaginationResponse, error) {
	authors, totalAuthors, err := s.Repository.GetAuthorsByFilter(models.Filter{
		Pagination: models.Pagination{
			Page:     int(req.Page),
			PageSize: int(req.PageSize),
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetAuthorsByFilter] Service error getting authors")
		return []dto.AuthorResponse{}, dto.PaginationResponse{}, err
	}

	responses := dto.BuildAuthorsResponse(authors)
	metadata := pagination.CalculatePaginationMetadata(req.Page, req.PageSize, totalAuthors)
	return responses, metadata, nil
}
