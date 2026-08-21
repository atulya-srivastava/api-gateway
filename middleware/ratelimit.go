package middleware

import (
	"net/http"
	"strings"
	"sync"
	"time"
)

type Client struct {
	Attempts    int
	LastAttempt time.Time
}

var (
	clients = make(map[string]Client)
	mu      sync.Mutex
)

const (
	MaxRequests = 10000
	Window      = 30 * time.Second
)

func RateLimit(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := strings.Split(r.RemoteAddr, ":")[0]

		mu.Lock()
		client := clients[ip]

		if time.Since(client.LastAttempt) > Window {
			client.Attempts = 0
			client.LastAttempt = time.Now()
		}

		client.Attempts++
		clients[ip] = client
		attempts := client.Attempts
		mu.Unlock()

		if attempts > MaxRequests {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}