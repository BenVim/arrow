package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
	http.ListenAndServe("0.0.0.0:8000", nil)
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(".....")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}