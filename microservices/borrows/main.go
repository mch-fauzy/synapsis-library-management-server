package main

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/synapsis-library-management-server/microservices/borrows/configs"
	"github.com/synapsis-library-management-server/microservices/borrows/handlers"
	"github.com/synapsis-library-management-server/microservices/borrows/repositories"
	"github.com/synapsis-library-management-server/microservices/borrows/routes"
	"github.com/synapsis-library-management-server/microservices/borrows/services"
)

func main() {
	// Initialize logger
	configs.InitLogger()

	// Initialize the PostgreSQL connection
	config := configs.Get()
	dbConn := configs.NewPostgreSqlConn(config)

	// Initialize repository, service, and handler layers
	repository := repositories.NewRepository(dbConn)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service)

	// Setup the router
	route := routes.SetupRouter(handler)

	// Start the server
	log.Info().Str("port", config.Server.Port).Msg("Starting up HTTP server")
	err := http.ListenAndServe(":"+config.Server.Port, route)
	if err != nil {
		log.Error().Err(err).Msg("Server failed to start")
	}
}
