package middlewares

import (
	"net/http"
	"strings"
	"time"
)

type clientInfo struct {
	Count     int
	ExpiresAt time.Time
}

var requestCounts = make(map[string]clientInfo)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userIP string
		if forwarded := strings.Split(r.Header.Get("x-forwarded-for"), ","); len(forwarded) > 0 {
			userIP = strings.TrimSpace(forwarded[0])
		} else {
			userIP = strings.Split(r.RemoteAddr, ":")[0]
		}
		// Implement rate limiting logic here using userIP without persisting data
		const limit = 50
		const window = time.Minute

		now := time.Now()
		ci, ok := requestCounts[userIP]
		if !ok || now.After(ci.ExpiresAt) {
			requestCounts[userIP] = clientInfo{
				Count:     1,
				ExpiresAt: now.Add(window),
			}
		} else {
			if ci.Count >= limit {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("Rate limit exceeded. Try again later."))
				return
			}
			ci.Count++
			requestCounts[userIP] = ci
		}

		next.ServeHTTP(w, r)
	})
}
