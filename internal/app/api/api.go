package api

import (
	"net/http"
	"simplewebserver/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Base API server instance description
type API struct {
	// UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	// Field for BD
	storage *storage.Storage
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/config, loggers, router, database connections etc...
func (api *API) Start() error {
	// Configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("Starting api server at port", api.config.BindAddr)

	// Configure mux (router)
	api.configureRouterField()

	// Configure storage
	if err := api.configStorageField(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
