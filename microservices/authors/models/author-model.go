package models

import (
	"time"

	"github.com/guregu/null"
	"gorm.io/gorm"
)

type authorDbFieldStruct struct {
	Id        string
	Name      string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
	DeletedAt string
	DeletedBy string
}

var AuthorDbField = authorDbFieldStruct{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
	DeletedAt: "deleted_at",
	DeletedBy: "deleted_by",
}

type Author struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy string
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy string
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy null.String
}

type AuthorPrimaryId struct {
	Id int `gorm:"primaryKey"`
}
