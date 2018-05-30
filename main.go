 package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	return r
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World!")
}