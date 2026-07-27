package main

import (
	"atulya/api-gateway/proxy"
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

	fmt.Fprintln(w,"Hello from user service")
}

func main(){

	userProxy := proxy.NewProxy("http://localhost/8001")

	http.Handle("/users", userProxy)

	log.Println("Backend running at port 8081")

	log.Fatal(http.ListenAndServe(":8081",nil))
}

