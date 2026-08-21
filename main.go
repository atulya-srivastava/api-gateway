package main

import (
	"atulya/api-gateway/healthcheck"
	loadbalencer "atulya/api-gateway/loadbalancer"
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

			healthyServers := []string{}

			results := make(chan Result,len(targetUrls))

			for _, server := range targetUrls {

				go func(server string){
					healthy := healthcheck.IsHealthy(server)

					results <- Result {
						serverName: server,
						healthy: healthy,
					}
					
				}(server)
				
			}

			for i:=0;i<len(targetUrls);i++{
				result := <- results
				if result.healthy {
					healthyServers = append(healthyServers,result.serverName)
				}
			}

			// If no backend is healthy, skip this route
			if len(healthyServers) == 0 {
				log.Println("No healthy servers for", route)
				continue
			}

			lb := loadbalencer.New(healthyServers)

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