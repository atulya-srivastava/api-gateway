package main

import (
	"atulya/api-gateway/proxy"
	"log"
	"net/http"
)

func main() {

	userProxy := proxy.NewProxy("http://localhost:8081")

	http.Handle("/users", userProxy)

	log.Println("Gateway running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}