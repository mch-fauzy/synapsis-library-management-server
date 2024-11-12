package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
)

type borrowDbFieldStruct struct {
	Id         string
	UserId     string
	BookId     string
	BorrowDate string
	ReturnDate string
	CreatedAt  string
	CreatedBy  string
	UpdatedAt  string
	UpdatedBy  string
	DeletedAt  string
	DeletedBy  string
}

var BorrowDbField = borrowDbFieldStruct{
	Id:         "id",
	UserId:     "user_id",
	BookId:     "book_id",
	BorrowDate: "borrow_date",
	ReturnDate: "return_date",
	CreatedAt:  "created_at",
	CreatedBy:  "created_by",
	UpdatedAt:  "updated_at",
	UpdatedBy:  "updated_by",
	DeletedAt:  "deleted_at",
	DeletedBy:  "deleted_by",
}

type Borrow struct {
	Id         int `gorm:"primaryKey"`
	UserId     uuid.UUID
	BookId     int
	BorrowDate time.Time
	ReturnDate null.Time
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	CreatedBy  string
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	UpdatedBy  string
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	DeletedBy  null.String
}

type BorrowPrimaryId struct {
	Id int `gorm:"primaryKey"`
}
