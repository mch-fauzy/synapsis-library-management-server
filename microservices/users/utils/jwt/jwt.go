package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/users/configs"
	"github.com/synapsis-library-management-server/microservices/users/models/dto"
)

func SignJwtToken(req dto.TokenPayload, tokenType string, expiryDuration time.Duration) (dto.LoginResponse, error) {
	config := configs.Get()

	expireTime := time.Now().Add(expiryDuration).Unix()

	// Create a new token with standard and custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": req.UserId,
		"email":  req.Email,
		"role":   req.Role,
		"exp":    expireTime,
	})

	// Sign the token with the provided secret key
	tokenString, err := token.SignedString([]byte(config.App.JwtAccessKey))
	if err != nil {
		log.Error().Err(err).Msg("[SignJWTToken] Failed to sign JWT Token")
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token:     tokenString,
		TokenType: tokenType,
		ExpiresIn: expireTime,
	}, nil
}
