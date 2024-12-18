package v1

import (
	"github.com/go-chi/chi"
	"github.com/synapsis-library-management-server/microservices/categories/handlers"
	"github.com/synapsis-library-management-server/microservices/categories/middlewares"
)

func V1Routes(r chi.Router, handler handlers.Handler) {
	r.Route("/", func(r chi.Router) {
		// Apply the authentication middleware
		r.Use(middlewares.AuthenticateToken)
		r.Get("/categories", handler.GetCategoriesByFilter)
		r.Get("/categories/{id}", handler.GetCategoryById)

		r.Group(func(r chi.Router) {
			r.Use(middlewares.AuthorizeAdmin)
			r.Post("/categories", handler.CreateCategory)
		})
	})
}
