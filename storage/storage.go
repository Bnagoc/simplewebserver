package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	config *Config
	// Database file descriptor
	db *sql.DB
}

// Storage constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open connection method
func (storage *Storage) Open() error {
	// Validate agrs (sql.Open())
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	// Verify a connection
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("DB connection created successfully!")
	return nil
}

// Close connection method
func (storage *Storage) Close() {
	storage.db.Close()
}
