package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
)

type userDbFieldStruct struct {
	Id        string
	Email     string
	Password  string
	Role      string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
	DeletedAt string
	DeletedBy string
}

var UserDbField = userDbFieldStruct{
	Id:        "id",
	Email:     "email",
	Password:  "password",
	Role:      "role",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
	DeletedAt: "deleted_at",
	DeletedBy: "deleted_by",
}

type User struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string
	Role      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreatedBy string
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedBy string
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy null.String
}
