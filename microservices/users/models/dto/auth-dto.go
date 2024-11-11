package dto

import (
	"github.com/gofrs/uuid"
	"github.com/synapsis-library-management-server/microservices/users/models"
	"github.com/synapsis-library-management-server/microservices/users/utils/failure"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"-"`
}

func (r RegisterRequest) Validate() error {
	if r.Email == "" {
		return failure.BadRequest("Email is required")
	}

	if r.Password == "" {
		return failure.BadRequest("Password is required")
	}

	if len(r.Password) < 8 {
		return failure.BadRequest("Password must be at least 8 characters")
	}

	return nil
}

func (r RegisterRequest) ToModel() models.User {
	id, _ := uuid.NewV4()
	return models.User{
		Id:        id,
		Email:     r.Email,
		Password:  r.Password,
		Role:      r.Role,
		CreatedBy: r.Email,
		UpdatedBy: r.Email,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequest) Validate() error {
	if r.Email == "" {
		return failure.BadRequest("Email is required")
	}

	if r.Password == "" {
		return failure.BadRequest("Password is required")
	}

	return nil
}

type LoginResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
	ExpiresIn int64  `json:"expiresIn"`
}

type TokenPayload struct {
	UserId string
	Email  string
	Role   string
}
