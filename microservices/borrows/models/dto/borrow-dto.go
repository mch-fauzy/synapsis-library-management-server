package dto

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/synapsis-library-management-server/microservices/borrows/models"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/constant"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/failure"
)

type CreateBorrowRequest struct {
	UserId uuid.UUID `json:"userId"`
	BookId int       `json:"bookId"`
	Email  string    `json:"-"`
}

func (r CreateBorrowRequest) Validate() error {
	if r.UserId == uuid.Nil {
		return failure.BadRequest("UserId is required")
	}

	if r.BookId <= 0 {
		return failure.BadRequest("BookId is required")
	}

	return nil
}

func (r CreateBorrowRequest) ToModel() models.Borrow {
	return models.Borrow{
		UserId:     r.UserId,
		BookId:     r.BookId,
		BorrowDate: time.Now(),
		CreatedBy:  r.Email,
		UpdatedBy:  r.Email,
	}
}

type GetBorrowsByFilterRequest struct {
	UserId   null.String `json:"userId"`
	Page     int64       `json:"page"`
	PageSize int64       `json:"pageSize"`
}

func (r GetBorrowsByFilterRequest) Validate() error {
	if r.Page <= 0 {
		r.Page = 1
	}

	if r.PageSize <= 0 {
		r.PageSize = 10
	}

	return nil
}

type BorrowResponse struct {
	Id         int         `json:"id"`
	UserId     uuid.UUID   `json:"userId"`
	BookId     int         `json:"bookId"`
	BorrowDate string      `json:"borrowDate"`
	ReturnDate null.String `json:"returnDate"`
	CreatedAt  string      `json:"createdAt"`
	CreatedBy  string      `json:"createdBy"`
	UpdatedAt  string      `json:"updatedAt"`
	UpdatedBy  string      `json:"updatedBy"`
}

func BuildBorrowsResponse(borrows []models.Borrow) []BorrowResponse {
	var responses []BorrowResponse
	for _, borrow := range borrows {

		returnDate := null.StringFrom(borrow.ReturnDate.Time.Format(constant.DateTimeUTCFormat))
		if !borrow.ReturnDate.Valid {
			returnDate = null.StringFromPtr(nil)
		}

		responses = append(responses, BorrowResponse{
			Id:         borrow.Id,
			UserId:     borrow.UserId,
			BookId:     borrow.BookId,
			BorrowDate: borrow.BorrowDate.Format(constant.DateTimeUTCFormat),
			ReturnDate: returnDate,
			CreatedAt:  borrow.CreatedAt.Format(constant.DateTimeUTCFormat),
			CreatedBy:  borrow.CreatedBy,
			UpdatedAt:  borrow.UpdatedAt.Format(constant.DateTimeUTCFormat),
			UpdatedBy:  borrow.UpdatedBy,
		})
	}
	return responses
}

type MarkBorrowAsReturnedByIdRequest struct {
	Id    int    `json:"-"`
	Email string `json:"-"`
}

func (r MarkBorrowAsReturnedByIdRequest) ToModel() models.Borrow {
	return models.Borrow{
		ReturnDate: null.TimeFrom(time.Now()),
		UpdatedBy:  r.Email,
	}
}
