package v1

import (
	"github.com/go-chi/chi"
	"github.com/synapsis-library-management-server/microservices/users/handlers"
)

func V1Routes(r chi.Router, handler handlers.Handler) {
	r.Post("/register", handler.RegisterUser)
	r.Post("/login", handler.Login)

	r.Route("/admin", func(r chi.Router) {
		r.Post("/register", handler.RegisterAdmin)
	})
}
