package routers

import (
	"mailverifier/api/handlers"
	"mailverifier/api/middlewares"

	"github.com/gorilla/mux"
)

func MailRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.ApplyCORS)
	router.Use(middlewares.HandleError)
	router.Use(middlewares.GenerateLogs)

	router.HandleFunc("/api/verify", handlers.VerifyMail).Methods("GET", "OPTIONS")

	return router
}
