package middlewares

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func ApplyCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var allowedOrigins = map[string]bool{
			"localhost:5173":    true,
			"mailverifier-xdps": true,
		}

		origin := r.Header.Get("origin")

		originUrl, err := url.Parse(origin)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Invalid origin data")
			return
		}
		if _, exists := allowedOrigins[originUrl.Host]; exists {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Headers", "content-type")
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Access forbidden")
			return
		}
	})
}
