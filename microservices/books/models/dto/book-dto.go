package dto

import (
	"github.com/synapsis-library-management-server/microservices/books/models"
	"github.com/synapsis-library-management-server/microservices/books/utils/constant"
	"github.com/synapsis-library-management-server/microservices/books/utils/failure"
)

type CreateBookRequest struct {
	Title         string `json:"title"`
	Stock         int    `json:"stock"`
	PublishedYear int    `json:"publishedYear"`
	Isbn          string `json:"isbn"`
	AuthorId      int    `json:"authorId"`
	CategoryId    int    `json:"categoryId"`
	Email         string `json:"-"`
}

func (r CreateBookRequest) Validate() error {
	if r.Title == "" {
		return failure.BadRequest("Title is required")
	}

	if r.Isbn == "" {
		return failure.BadRequest("ISBN is required")
	}

	if r.PublishedYear <= 0 {
		return failure.BadRequest("Published Year must be greater than 0")
	}

	if r.AuthorId <= 0 {
		return failure.BadRequest("Author ID is required")
	}

	if r.CategoryId <= 0 {
		return failure.BadRequest("Category ID is required")
	}

	return nil
}

func (r CreateBookRequest) ToModel() models.Book {
	return models.Book{
		Title:         r.Title,
		Stock:         r.Stock,
		PublishedYear: r.PublishedYear,
		Isbn:          r.Isbn,
		AuthorId:      r.AuthorId,
		CategoryId:    r.CategoryId,
		CreatedBy:     r.Email,
		UpdatedBy:     r.Email,
	}
}

type GetBooksByFilterRequest struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

func (r GetBooksByFilterRequest) Validate() error {
	if r.Page <= 0 {
		r.Page = 1
	}

	if r.PageSize <= 0 {
		r.PageSize = 10
	}

	return nil
}

type GetBooksByFilterResponse struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Stock         int    `json:"stock"`
	PublishedYear int    `json:"publishedYear"`
	Isbn          string `json:"isbn"`
	AuthorId      int    `json:"authorId"`
	CategoryId    int    `json:"categoryId"`
	CreatedAt     string `json:"createdAt"`
	CreatedBy     string `json:"createdBy"`
	UpdatedAt     string `json:"updatedAt"`
	UpdatedBy     string `json:"updatedBy"`
}

func BuildGetBooksByFilterResponse(books []models.Book) []GetBooksByFilterResponse {
	var responses []GetBooksByFilterResponse
	for _, book := range books {
		responses = append(responses, GetBooksByFilterResponse{
			Id:            book.Id,
			Title:         book.Title,
			Stock:         book.Stock,
			PublishedYear: book.PublishedYear,
			Isbn:          book.Isbn,
			AuthorId:      book.AuthorId,
			CategoryId:    book.CategoryId,
			CreatedAt:     book.CreatedAt.Format(constant.DateTimeUTCFormat),
			CreatedBy:     book.CreatedBy,
			UpdatedAt:     book.UpdatedAt.Format(constant.DateTimeUTCFormat),
			UpdatedBy:     book.UpdatedBy,
		})
	}
	return responses
}
