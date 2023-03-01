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
func (s *Storage) Open() error {
	// Validate agrs (sql.Open())
	db, err := sql.Open("postgres", s.config.DatabaseURI)
	if err != nil {
		return err
	}
	// Verify a connection
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("DB connection created successfully!")
	return nil
}

// Close connection method
func (s *Storage) Close() {
	s.db.Close()
}

// Public repo for Article
func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		storage: s,
	}
	return s.articleRepository
}

// Public repo for User
func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return s.userRepository
}
