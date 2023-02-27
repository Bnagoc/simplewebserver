package storage

import (
	"fmt"
	"log"
	"simplewebserver/internal/app/models"
)

// Instance of Article repository (model interface)
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "articles"
)

// Add Article in DB
func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUSE ($1, $2, $3) RETURNING ID", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

// Delete Article from DB
func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	article, ok, err := ar.FindArticleById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableArticle)
		_, err := ar.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}
	return article, nil
}

// Select article by ID
func (ar *ArticleRepository) FindArticleById(id int) (*models.Article, bool, error) {
	articles, err := ar.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var articleFound *models.Article
	for _, a := range articles {
		if a.ID == id {
			articleFound = a
			founded = true
			break
		}
	}
	return articleFound, founded, nil
}

// Select all Articles from DB
func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Where to read
	articles := make([]*models.Article, 0)
	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}

// Update the article
