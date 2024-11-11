package dto

import "github.com/dgrijalva/jwt-go"

type TokenPayload struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
