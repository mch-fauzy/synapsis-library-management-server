package services

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/borrows/models"
	"github.com/synapsis-library-management-server/microservices/borrows/models/dto"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/failure"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/pagination"
	"gorm.io/gorm"
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

	filterFields := []models.FilterField{}

	if req.UserId.Valid && req.UserId.String != "" {
		filterFields = append(filterFields, models.FilterField{
			Field:    models.BorrowDbField.UserId,
			Operator: models.OperatorEqual,
			Value:    req.UserId.String,
		})
	}

	borrows, totalBorrows, err := s.Repository.GetBorrowsByFilter(models.Filter{
		FilterFields: filterFields,
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

func (s *Service) MarkBorrowAsReturnedById(req dto.MarkBorrowAsReturnedByIdRequest) (string, error) {
	message := "Failed"

	borrow, err := s.Repository.GetBorrowById(models.BorrowPrimaryId{Id: req.Id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = failure.NotFound("Borrow not found")
			return message, err
		}

		log.Error().Err(err).Msg("[MarkBorrowAsReturnedById] Service error retrieving borrow by id")
		return message, err
	}

	borrowToUpdate := req.ToModel()
	err = s.Repository.UpdateBorrow(models.BorrowPrimaryId{Id: borrow.Id}, &borrowToUpdate)
	if err != nil {
		log.Error().Err(err).Msg("[MarkBorrowAsReturnedById] Service error updating borrow as returned")
		return message, err
	}

	message = "Success"
	return message, nil
}
