package configs

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	maxIdleConnection = 10
	maxOpenConnection = 10
)

// PostgreSqlConn wraps a single GORM database connection for both read and write operations
type PostgreSqlConn struct {
	Db *gorm.DB
}

// NewPostgreSqlConn is the constructor for PostgreSqlConn
func NewPostgreSqlConn(config *Config) *PostgreSqlConn {
	return &PostgreSqlConn{
		Db: CreatePostgreSqlConn(*config),
	}
}

// CreatePostgreSqlConn creates a GORM database connection for both read and write access
func CreatePostgreSqlConn(config Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DB.PostgreSQL.Host, // Assuming read and write are on the same server
		config.DB.PostgreSQL.Port,
		config.DB.PostgreSQL.Username,
		config.DB.PostgreSQL.Password,
		config.DB.PostgreSQL.Name,
		"disable",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Failed to connect to database")
	} else {
		log.
			Info().
			Str("host", config.DB.PostgreSQL.Host).
			Str("port", config.DB.PostgreSQL.Port).
			Str("dbName", config.DB.PostgreSQL.Name).
			Msg("Connected to database")
	}

	// Set connection pool parameters
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get the underlying SQL Db")
	}

	sqlDb.SetMaxIdleConns(maxIdleConnection)
	sqlDb.SetMaxOpenConns(maxOpenConnection)

	return db
}
