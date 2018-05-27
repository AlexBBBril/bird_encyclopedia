package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hellow World!")
}