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
	// Subfield for repo interface (model user)
	userRepository *UserRepository
	// Subfield for repo interface (model article)
	articleRepository *ArticleRepository
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

// Public repo for Article
func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		storage: s,
	}
	return nil
}

// Public repo for User
func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return nil
}
