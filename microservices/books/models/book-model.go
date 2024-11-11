package models

import (
	"time"

	"github.com/guregu/null"
	"gorm.io/gorm"
)

type bookDbFieldStruct struct {
	Id            string
	Title         string
	Stock         string
	PublishedYear string
	Isbn          string
	AuthorId      string
	CategoryId    string
	CreatedAt     string
	CreatedBy     string
	UpdatedAt     string
	UpdatedBy     string
	DeletedAt     string
	DeletedBy     string
}

var BookDbField = bookDbFieldStruct{
	Id:            "id",
	Title:         "title",
	Stock:         "stock",
	PublishedYear: "published_year",
	Isbn:          "isbn",
	AuthorId:      "author_id",
	CategoryId:    "category_id",
	CreatedAt:     "created_at",
	CreatedBy:     "created_by",
	UpdatedAt:     "updated_at",
	UpdatedBy:     "updated_by",
	DeletedAt:     "deleted_at",
	DeletedBy:     "deleted_by",
}

type Book struct {
	Id            int `gorm:"primaryKey"`
	Title         string
	Stock         int
	PublishedYear int
	Isbn          string `gorm:"uniqueIndex"`
	AuthorId      int
	CategoryId    int
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	CreatedBy     string
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	UpdatedBy     string
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	DeletedBy     null.String
}
