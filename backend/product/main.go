package main

import (
	"fmt"
	"log"
	"net/http"
)

func productHandler( w http.ResponseWriter, r *http.Request ){

	fmt.Fprintln(w,"hello there from the products service");
}

func main(){

	http.HandleFunc("/products",productHandler);

	log.Println("Product service is listening on port")

	http.ListenAndServe(":8082",nil)

}