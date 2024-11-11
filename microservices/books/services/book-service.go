package services

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/books/models"
	"github.com/synapsis-library-management-server/microservices/books/models/dto"
	"github.com/synapsis-library-management-server/microservices/books/utils/failure"
	"github.com/synapsis-library-management-server/microservices/books/utils/pagination"
)

func (s *Service) CreateBook(req dto.CreateBookRequest) (string, error) {
	message := "Failed"

	_, totalBooks, err := s.Repository.GetBooksByFilter(models.Filter{
		FilterFields: []models.FilterField{
			{
				Field:    models.BookDbField.Isbn,
				Operator: models.OperatorEqual,
				Value:    req.Isbn,
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[CreateBook] Service error getting books")
		return message, err
	}

	if totalBooks > 0 {
		err = failure.Conflict("Book with this ISBN already exists")
		return message, err
	}

	// TODO: get author id from authors microservice
	// TODO: get category id from cateories microservice

	book := req.ToModel()
	err = s.Repository.CreateBook(&book)
	if err != nil {
		log.Error().Err(err).Msg("[CreateBook] Service error creating book")
		return message, err
	}

	message = "Success"
	return message, nil
}

func (s *Service) GetBooksByFilter(req dto.GetBooksByFilterRequest) ([]dto.GetBooksByFilterResponse, dto.PaginationResponse, error) {
	books, totalBooks, err := s.Repository.GetBooksByFilter(models.Filter{
		Pagination: models.Pagination{
			Page:     int(req.Page),
			PageSize: int(req.PageSize),
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetBooksByFilter] Service error getting books")
		return []dto.GetBooksByFilterResponse{}, dto.PaginationResponse{}, err
	}

	responses := dto.BuildGetBooksByFilterResponse(books)
	metadata := pagination.CalculatePaginationMetadata(req.Page, req.PageSize, totalBooks)
	return responses, metadata, nil
}
