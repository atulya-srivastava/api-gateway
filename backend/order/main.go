package main

import (
	"fmt"
	"log"
	"net/http"
)

func orderHandler( w http.ResponseWriter, r *http.Request ){
	log.Print("this is the method", r.Method);
	log.Print("this is the url", r.URL.Path)
	fmt.Fprintln(w,"hello there from order service");
}


func healthHandler(w http.ResponseWriter, r * http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok the server order is responding")
}

func main(){

	http.HandleFunc("/orders",orderHandler);
	http.HandleFunc("/health", healthHandler);

	log.Println("Order service is listening on port 8083")

	http.ListenAndServe(":8083",nil)

}