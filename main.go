package main

import (
	loadbalencer "atulya/api-gateway/loadbalancer"
	"atulya/api-gateway/middleware"
	"atulya/api-gateway/healthcheck"
	"atulya/api-gateway/services"
	"log"
	"net/http"
)

	func main() {

		// userProxy := proxy.NewProxy("http://localhost:8081")
		// productProxy := proxy.NewProxy("http://localhost:8082")
		// orderProxy := proxy.NewProxy("http://localhost:8083")

		// http.Handle("/users", userProxy)
		// http.Handle("/products", productProxy)
		// http.Handle("/orders",orderProxy)

		for route, targetUrls := range services.Routes{

			healthyServers := []string{}

			for _, server := range targetUrls {

				// Check whether this backend is alive
				if healthcheck.IsHealthy(server) {
					healthyServers = append(healthyServers, server)
				}
			}

			// If no backend is healthy, skip this route
			if len(healthyServers) == 0 {
				log.Println("No healthy servers for", route)
				continue
			}

			lb := loadbalencer.New(targetUrls)

			http.Handle(
				route,
				middleware.Authentication(
					middleware.RateLimit(
					middleware.Logging(lb),
				),
			),
		)
		
	}
	
	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}