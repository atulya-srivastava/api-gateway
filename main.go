package main

import (
	"atulya/api-gateway/healthcheck"
	"atulya/api-gateway/loadbalancer"
	"atulya/api-gateway/middleware"
	"atulya/api-gateway/services"
	"log"
	"net/http"
)

type Result struct{
	serverName string
	healthy bool
}

	func main() {

		// userProxy := proxy.NewProxy("http://localhost:8081")
		// productProxy := proxy.NewProxy("http://localhost:8082")
		// orderProxy := proxy.NewProxy("http://localhost:8083")

		// http.Handle("/users", userProxy)
		// http.Handle("/products", productProxy)
		// http.Handle("/orders",orderProxy)

		for route, targetUrls := range services.Routes{

			healthyServers := healthcheck.CheckAll(targetUrls)

			// If no backend is healthy, skip this route
			if len(healthyServers) == 0 {
				log.Println("No healthy servers for", route)
				continue
			}

			lb := loadbalancer.New(healthyServers)

			http.Handle(
				route,
				middleware.Authentication(
					middleware.RateLimit(
					middleware.Logging(lb),
				),
			),
		)
		
		go healthcheck.StartMonitor(targetUrls,lb.UpdateServers)
	}
	
	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}