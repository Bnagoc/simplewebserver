package api

import (
	"net/http"
	"simplewebserver/storage"

	"github.com/sirupsen/logrus"
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
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there!"))
	})
}

// Configurating DB repository (storage API)
func (a *API) configStorageField() error {
	storage := storage.New(a.config.Storage)
	// Connect to DB if possible, or throw error
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
