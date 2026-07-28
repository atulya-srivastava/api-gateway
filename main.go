	package main

	import (
		"atulya/api-gateway/proxy"
		"log"
		"net/http"
		"atulya/api-gateway/services"
	)

	func main() {

		// userProxy := proxy.NewProxy("http://localhost:8081")
		// productProxy := proxy.NewProxy("http://localhost:8082")
		// orderProxy := proxy.NewProxy("http://localhost:8083")

		// http.Handle("/users", userProxy)
		// http.Handle("/products", productProxy)
		// http.Handle("/orders",orderProxy)

		for route, targetUrl := range services.Routes{

			urlProxy := proxy.NewProxy(targetUrl)

			http.Handle(route,urlProxy)
		}

		log.Println("Gateway running on :8080")

		log.Fatal(http.ListenAndServe(":8080", nil))
	}