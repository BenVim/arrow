package server

import (
	"net/http"
	"fmt"
)

type httpServer struct {

}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(".....")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}


func StartServer() {
	http.HandleFunc("/", handler) // each request calls handler
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
	http.ListenAndServe("0.0.0.0:8000", nil)
}


