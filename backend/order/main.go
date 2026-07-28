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

func main(){

	http.HandleFunc("/orders",orderHandler);

	log.Println("Order service is listening on port 8083")

	http.ListenAndServe(":8083",nil)

}