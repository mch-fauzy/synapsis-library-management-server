package models

import (
	"time"

	"github.com/guregu/null"
	"gorm.io/gorm"
)

type categoryDbFieldStruct struct {
	Id        string
	Name      string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
	DeletedAt string
	DeletedBy string
}

var CategoryDbField = categoryDbFieldStruct{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
	DeletedAt: "deleted_at",
	DeletedBy: "deleted_by",
}

type Category struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy string
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy string
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy null.String
}
