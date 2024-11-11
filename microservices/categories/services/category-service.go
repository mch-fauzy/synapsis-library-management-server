package services

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/categories/models"
	"github.com/synapsis-library-management-server/microservices/categories/models/dto"
	"github.com/synapsis-library-management-server/microservices/categories/utils/failure"
	"github.com/synapsis-library-management-server/microservices/categories/utils/pagination"
	"gorm.io/gorm"
)

func (s *Service) CreateCategory(req dto.CreateCategoryRequest) (string, error) {
	message := "Failed"

	_, totalCategories, err := s.Repository.GetCategoriesByFilter(models.Filter{
		FilterFields: []models.FilterField{
			{
				Field:    models.CategoryDbField.Name,
				Operator: models.OperatorEqual,
				Value:    req.Name,
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[CreateCategory] Service error getting categoris")
		return message, err
	}

	if totalCategories > 0 {
		err = failure.Conflict("Category with this name already exists")
		return message, err
	}

	category := req.ToModel()
	err = s.Repository.CreateCategory(&category)
	if err != nil {
		log.Error().Err(err).Msg("[CreateCategories] Service error creating category")
		return message, err
	}

	message = "Success"
	return message, nil
}

func (s *Service) GetCategoryById(req dto.GetCategoryByIdRequest) (dto.CategoryResponse, error) {

	category, err := s.Repository.GetCategoryById(models.CategoryPrimaryId{Id: req.Id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = failure.NotFound("Category not found")
			return dto.CategoryResponse{}, err
		}

		log.Error().Err(err).Msg("[GetCategoryById] Service error retrieving category by id")
		return dto.CategoryResponse{}, err
	}

	response := dto.NewCategoryResponse(category)
	return response, nil
}

func (s *Service) GetCategoriesByFilter(req dto.GetCategoriesByFilterRequest) ([]dto.CategoryResponse, dto.PaginationResponse, error) {
	categories, totalCategories, err := s.Repository.GetCategoriesByFilter(models.Filter{
		Pagination: models.Pagination{
			Page:     int(req.Page),
			PageSize: int(req.PageSize),
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetCategoriesByFilter] Service error getting categories")
		return []dto.CategoryResponse{}, dto.PaginationResponse{}, err
	}

	responses := dto.BuildCategoriesResponse(categories)
	metadata := pagination.CalculatePaginationMetadata(req.Page, req.PageSize, totalCategories)
	return responses, metadata, nil
}
