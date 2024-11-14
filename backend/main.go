package main

import (
	"mailverifier/api/routers"
	"net/http"
)

func main() {
	router := routers.MailRouter()
	http.ListenAndServe(":8080", router)
}
