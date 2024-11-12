package v1

import (
	"github.com/go-chi/chi"
	"github.com/synapsis-library-management-server/microservices/borrows/handlers"
	"github.com/synapsis-library-management-server/microservices/borrows/middlewares"
)

func V1Routes(r chi.Router, handler handlers.Handler) {
	r.Route("/", func(r chi.Router) {
		r.Use(middlewares.AuthenticateToken)

		r.Group(func(r chi.Router) {
			r.Use(middlewares.AuthorizeAdmin)
			r.Post("/borrows", handler.CreateBorrow)
			r.Get("/borrows", handler.GetBorrowsByFilter)
			r.Patch("/borrows/{id}", handler.MarkBorrowAsReturnedById)
		})
	})
}
