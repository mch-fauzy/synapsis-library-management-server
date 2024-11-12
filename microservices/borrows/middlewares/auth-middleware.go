package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/borrows/configs"
	"github.com/synapsis-library-management-server/microservices/borrows/models/dto"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/constant"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/failure"
	"github.com/synapsis-library-management-server/microservices/borrows/utils/response"
)

// AuthenticateToken middleware validates the JWT token in the request header
func AuthenticateToken(next http.Handler) http.Handler {
	config := configs.Get()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.WithMessage(w, http.StatusUnauthorized, "Missing token")
			return
		}

		// Extract token from "Bearer <Token>" and parse the JWT token
		tokenString := strings.Split(authHeader, " ")[1]
		tokenClaims := &dto.TokenPayload{}
		token, err := jwt.ParseWithClaims(tokenString, tokenClaims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := failure.InternalError("Unexpected signing method")
				return nil, err
			}

			return []byte(config.App.JwtAccessKey), nil
		})
		if err != nil || !token.Valid {
			log.Error().Err(err).Msg("[AuthenticateToken] Token invalid")
			response.WithError(w, err)
			return
		}

		// Extract the claims from the token
		claims, ok := token.Claims.(*dto.TokenPayload) // Type Assertion to TokenPayload type
		if !ok {
			response.WithMessage(w, http.StatusUnauthorized, "Invalid token payload")
			return
		}

		// Check if userId, email, and role are present
		if claims.UserId == "" || claims.Email == "" || claims.Role == "" {
			response.WithMessage(w, http.StatusUnauthorized, "Incomplete token payload")
			return
		}

		// Set the decoded token details in the request headers
		r.Header.Set(constant.UserIdHeader, claims.UserId)
		r.Header.Set(constant.EmailHeader, claims.Email)
		r.Header.Set(constant.RoleHeader, claims.Role)

		next.ServeHTTP(w, r)
	})
}

// AuthorizeAdmin middleware checks if the user is an admin
func AuthorizeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get(constant.RoleHeader)
		if role != constant.RoleAdmin {
			response.WithMessage(w, http.StatusForbidden, "Admin access only")
			return
		}

		next.ServeHTTP(w, r)
	})
}
