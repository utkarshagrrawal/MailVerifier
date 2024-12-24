package main

import (
	"mailverifier/api/routers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/api/v1").Handler(http.StripPrefix("/api/v1", routers.MailRouter()))

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}

	http.ListenAndServe(port, router)
}
