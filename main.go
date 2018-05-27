 package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World!")
}