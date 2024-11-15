package handlers

import (
	"encoding/json"
	"mailverifier/api/services"
	"net/http"
	"strings"
)

func VerifyMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	values := r.URL.Query()
	email := values.Get("email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Email not received")
		return
	}
	emailComponents := strings.Split(email, "@")
	if len(emailComponents) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid email received")
		return
	}
	mxValidity := services.ValidateMXRecord(emailComponents[1])
	if mxValidity != "Domain available" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(mxValidity)
		return
	}
	disposability := services.CheckDisposableDomain(emailComponents[1])
	if disposability != "Domain is not disposable" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(disposability)
		return
	}
	deliverability := services.CheckMailDeliverable(email)
	if deliverability != "Mail address is deliverable" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(deliverability)
		return
	}
	json.NewEncoder(w).Encode("Email is available")
}
