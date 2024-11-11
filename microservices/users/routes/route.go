package routes

import (
	"github.com/go-chi/chi"
	"github.com/synapsis-library-management-server/microservices/users/handlers"
	v1 "github.com/synapsis-library-management-server/microservices/users/routes/v1"
)

func SetupRouter(handler handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	// r.Use(middleware.Logger)      // Optional: logging middleware
	// r.Use(AuthMiddleware)         // Apply JWT authentication middleware

	r.Route("/v1", func(r chi.Router) {
		v1.V1Routes(r, handler)
	})

	return r
}
