package main

import (
	"atulya/api-gateway/healthcheck"
	"atulya/api-gateway/loadbalancer"
	"atulya/api-gateway/middleware"
	"atulya/api-gateway/services"
	"log"
	"net/http"
)

	func main() {

		for route, targetUrls := range services.Routes{

			healthyServers := healthcheck.CheckAll(targetUrls)
			
			if len(healthyServers) == 0 {
				log.Println("No healthy servers for", route)
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
		
		go healthcheck.StartMonitor(route,targetUrls,lb.UpdateServers)
	}
	
	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}