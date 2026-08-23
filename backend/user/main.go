package main

import (
	"fmt"
	"log"
	"net/http"
)

func usersHandler(w http.ResponseWriter, r *http.Request){

	// log.Println("received at backend meth ",r.Method," url ",r.URL.Path)
	log.Println("Method:", r.Method)
	log.Println("Path:", r.URL.Path)
	log.Println("Headers:", r.Header)
	log.Println("Host:", r.Host)
	log.Println("Remote Address:", r.RemoteAddr)

	fmt.Fprintln(w,"Hello from user service 1")
}

func healthHandler(w http.ResponseWriter, r * http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok the server user is responding")
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/health", healthHandler)
	
	log.Println("Backend running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

