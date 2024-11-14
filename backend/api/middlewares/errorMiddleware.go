package middlewares

import (
	"log"
	"net/http"
)

func HandleError(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rescue := recover(); rescue != nil {
				log.Printf("%s request to: %s failed due to panic: %s", r.Method, r.URL.Path, rescue)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
