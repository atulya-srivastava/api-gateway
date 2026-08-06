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

	fmt.Fprintln(w,"Hello from user service 2")
}

func main() {
	http.HandleFunc("/users", usersHandler)

	log.Println("Backend running on :8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}

