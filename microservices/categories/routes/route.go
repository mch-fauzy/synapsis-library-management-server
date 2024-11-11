package routes

import (
	"github.com/go-chi/chi"
	"github.com/synapsis-library-management-server/microservices/categories/handlers"
	v1 "github.com/synapsis-library-management-server/microservices/categories/routes/v1"
)

func SetupRouter(handler handlers.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		v1.V1Routes(r, handler)
	})

	return r
}
