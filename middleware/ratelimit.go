package middleware

import (
	"net/http"
	"strings"
	"time"
)

type Client struct {
	Attempts    int
	LastAttempt time.Time
}

var clients = make(map[string]Client)

const (
	MaxRequests = 5
	Window      = 30 * time.Second
)

func RateLimit(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := strings.Split(r.RemoteAddr, ":")[0]

		client := clients[ip]

		if time.Since(client.LastAttempt) > Window {
			client.Attempts = 0
			client.LastAttempt = time.Now()
		}

		client.Attempts++

		clients[ip] = client

		if client.Attempts > MaxRequests {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}