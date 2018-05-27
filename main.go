 package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewROuter()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hellow World!")
}