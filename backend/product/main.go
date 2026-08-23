package main

import (
	"fmt"
	"log"
	"net/http"
)

func productHandler( w http.ResponseWriter, r *http.Request ){

	fmt.Fprintln(w,"hello there from the products service");
}

func healthHandler(w http.ResponseWriter, r * http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok the server is responding")
}

func main(){

	http.HandleFunc("/products",productHandler);
	http.HandleFunc("/health", healthHandler);

	log.Println("Product service is listening on port")

	http.ListenAndServe(":8082",nil)

}