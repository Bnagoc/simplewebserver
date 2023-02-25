package api

import (
	"net/http"

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
