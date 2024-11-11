package services

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/users/models"
	"github.com/synapsis-library-management-server/microservices/users/models/dto"
	"github.com/synapsis-library-management-server/microservices/users/utils/failure"
	"github.com/synapsis-library-management-server/microservices/users/utils/jwt"
	"github.com/synapsis-library-management-server/microservices/users/utils/password"
)

func (s *Service) Register(req dto.RegisterRequest) (string, error) {
	message := "Failed"

	// Check if email already exists
	_, totalUsers, err := s.Repository.GetUsersByFilter(models.Filter{
		FilterFields: []models.FilterField{
			{
				Field:    models.UserDbField.Email,
				Operator: models.OperatorEqual,
				Value:    req.Email,
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[Register] Service error getting users")
		return message, err
	}

	if totalUsers > 0 {
		err = failure.Conflict("Account with this email already exists")
		return message, err
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		log.Error().Err(err).Msg("[Register] Service error hashing password")
		return message, err
	}

	req.Password = hashedPassword
	user := req.ToModel()
	err = s.Repository.CreateUser(&user)
	if err != nil {
		log.Error().Err(err).Msg("[Register] Service error creating user")
		return message, err
	}

	message = "Success"
	return message, nil
}

func (s *Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	users, totalUsers, err := s.Repository.GetUsersByFilter(models.Filter{
		FilterFields: []models.FilterField{
			{
				Field:    models.UserDbField.Email,
				Operator: models.OperatorEqual,
				Value:    req.Email,
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[Login] Service error getting users")
		return dto.LoginResponse{}, err
	}

	if totalUsers == 0 {
		err = failure.Unauthorized("Invalid credentials")
		return dto.LoginResponse{}, err
	}

	err = password.ComparePassword(req.Password, users[0].Password)
	if err != nil {
		log.Error().Err(err).Msg("[Login] Service error comparing password")
		err = failure.Unauthorized("Invalid credentials")
		return dto.LoginResponse{}, err
	}

	response, err := jwt.SignJwtToken(dto.TokenPayload{
		UserId: users[0].Id.String(),
		Email:  users[0].Email,
		Role:   users[0].Role,
	},
		"Bearer",
		time.Hour,
	)
	if err != nil {
		log.Error().Err(err).Msg("[Login] Service error signing jwt token")
		err = failure.InternalError("Failed to login user")
		return dto.LoginResponse{}, err
	}

	return response, nil
}
