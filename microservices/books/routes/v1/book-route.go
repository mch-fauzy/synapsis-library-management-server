package v1

import (
	"github.com/go-chi/chi"
	"github.com/synapsis-library-management-server/microservices/books/handlers"
	"github.com/synapsis-library-management-server/microservices/books/middlewares"
)

func V1Routes(r chi.Router, handler handlers.Handler) {
	r.Route("/", func(r chi.Router) {
		// Apply the authentication middleware
		r.Use(middlewares.AuthenticateToken)
		r.Get("/books", handler.GetBooksByFilter)

		r.Group(func(r chi.Router) {
			r.Use(middlewares.AuthorizeAdmin)
			r.Post("/books", handler.CreateBook)
		})
	})
}
