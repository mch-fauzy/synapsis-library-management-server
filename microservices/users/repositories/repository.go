package repositories

import (
	"github.com/synapsis-library-management-server/microservices/users/configs"
)

type Repository struct {
	PostgreSqlConn *configs.PostgreSqlConn
}

// NewRepository is the constructor for Repository
func NewRepository(db *configs.PostgreSqlConn) *Repository {
	return &Repository{
		PostgreSqlConn: db,
	}
}
