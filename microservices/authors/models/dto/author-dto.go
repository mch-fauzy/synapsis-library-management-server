package dto

import (
	"github.com/synapsis-library-management-server/microservices/authors/models"
	"github.com/synapsis-library-management-server/microservices/authors/utils/constant"
	"github.com/synapsis-library-management-server/microservices/authors/utils/failure"
)

type CreateAuthorRequest struct {
	Name  string `json:"name"`
	Email string `json:"-"`
}

func (r CreateAuthorRequest) Validate() error {
	if r.Name == "" {
		return failure.BadRequest("Name is required")
	}

	return nil
}

func (r CreateAuthorRequest) ToModel() models.Author {
	return models.Author{
		Name:      r.Name,
		CreatedBy: r.Email,
		UpdatedBy: r.Email,
	}
}

type GetAuthorsByFilterRequest struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

func (r GetAuthorsByFilterRequest) Validate() error {
	if r.Page <= 0 {
		r.Page = 1
	}

	if r.PageSize <= 0 {
		r.PageSize = 10
	}

	return nil
}

type GetAuthorByIdRequest struct {
	Id int `json:"-"`
}

type AuthorResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt string `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
}

func NewAuthorResponse(author models.Author) AuthorResponse {
	return AuthorResponse{
		Id:        author.Id,
		Name:      author.Name,
		CreatedAt: author.CreatedAt.Format(constant.DateTimeUTCFormat),
		CreatedBy: author.CreatedBy,
		UpdatedAt: author.UpdatedAt.Format(constant.DateTimeUTCFormat),
		UpdatedBy: author.UpdatedBy,
	}
}

func BuildAuthorsResponse(authors []models.Author) []AuthorResponse {
	var responses []AuthorResponse
	for _, author := range authors {
		responses = append(responses, AuthorResponse{
			Id:        author.Id,
			Name:      author.Name,
			CreatedAt: author.CreatedAt.Format(constant.DateTimeUTCFormat),
			CreatedBy: author.CreatedBy,
			UpdatedAt: author.UpdatedAt.Format(constant.DateTimeUTCFormat),
			UpdatedBy: author.UpdatedBy,
		})
	}
	return responses
}
