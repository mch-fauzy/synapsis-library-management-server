package services

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/borrows/models"
	"github.com/synapsis-library-management-server/microservices/borrows/models/dto"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/pagination"
)

func (s *Service) CreateBorrow(req dto.CreateBorrowRequest) (string, error) {
	message := "Failed"

	// TODO: get user id from users microservice, exist or not
	// TODO: get book id from books microservice, exist or not

	borrow := req.ToModel()
	err := s.Repository.CreateBorrow(&borrow)
	if err != nil {
		log.Error().Err(err).Msg("[CreateBorrow] Service error creating borrow")
		return message, err
	}

	message = "Success"
	return message, nil
}

func (s *Service) GetBorrowsByFilter(req dto.GetBorrowsByFilterRequest) ([]dto.BorrowResponse, dto.PaginationResponse, error) {
	borrows, totalBorrows, err := s.Repository.GetBorrowsByFilter(models.Filter{
		Pagination: models.Pagination{
			Page:     int(req.Page),
			PageSize: int(req.PageSize),
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetBorrowsByFilter] Service error getting borrows")
		return []dto.BorrowResponse{}, dto.PaginationResponse{}, err
	}

	responses := dto.BuildBorrowsResponse(borrows)
	metadata := pagination.CalculatePaginationMetadata(req.Page, req.PageSize, totalBorrows)
	return responses, metadata, nil
}