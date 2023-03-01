package api

import (
	"simplewebserver/storage"

	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

// Configurating API instance (logger field in API struct)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Configurating router (router field in API struct)
func (a *API) configureRouterField() {
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
	// New pair for auth
	a.router.HandleFunc(prefix+"/user/auth", a.PostToAuth).Methods("POST")
}

// Configurating DB repository (storage API)
func (a *API) configStorageField() error {
	st := storage.New(a.config.Storage)
	// Connect to DB if possible, or throw error
	if err := st.Open(); err != nil {
		return err
	}
	a.storage = st
	return nil
}
