package controllers

import "github.com/gorilla/mux"

const prefix = "/api"

var router *mux.Router

func NewRouter() *mux.Router {
	router = mux.NewRouter()
	router = router.PathPrefix(prefix).Subrouter()

	SetupHealthCheck(router)
	return router
}

func SetupHealthCheck(router *mux.Router) {
	router.HandleFunc("/health", HealthCheck).Methods("GET")
}
