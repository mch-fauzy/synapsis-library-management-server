package dto

import (
	"github.com/synapsis-library-management-server/microservices/categories/models"
	"github.com/synapsis-library-management-server/microservices/categories/utils/constant"
	"github.com/synapsis-library-management-server/microservices/categories/utils/failure"
)

type CreateCategoryRequest struct {
	Name  string `json:"name"`
	Email string `json:"-"`
}

func (r CreateCategoryRequest) Validate() error {
	if r.Name == "" {
		return failure.BadRequest("Name is required")
	}

	return nil
}

func (r CreateCategoryRequest) ToModel() models.Category {
	return models.Category{
		Name:      r.Name,
		CreatedBy: r.Email,
		UpdatedBy: r.Email,
	}
}

type GetCategoriesByFilterRequest struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

func (r GetCategoriesByFilterRequest) Validate() error {
	if r.Page <= 0 {
		r.Page = 1
	}

	if r.PageSize <= 0 {
		r.PageSize = 10
	}

	return nil
}

type GetCategoriesByFilterResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt string `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
}

func BuildGetCategoriesByFilterResponse(categories []models.Category) []GetCategoriesByFilterResponse {
	var responses []GetCategoriesByFilterResponse
	for _, category := range categories {
		responses = append(responses, GetCategoriesByFilterResponse{
			Id:        category.Id,
			Name:      category.Name,
			CreatedAt: category.CreatedAt.Format(constant.DateTimeUTCFormat),
			CreatedBy: category.CreatedBy,
			UpdatedAt: category.UpdatedAt.Format(constant.DateTimeUTCFormat),
			UpdatedBy: category.UpdatedBy,
		})
	}
	return responses
}
