package main

import (
	"fmt"
	"net/http"
)

func main() {

	var handler http.ServeMux
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/rocket":
			rktHandler(w, r)
		case "/":
			if r.Method != "GET" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Method %s not supported", r.Method)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Welcome to Go RPC Rockets MicroService")
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Invalid Route: Please visit /rocket")
			return

		}

	})

	server := http.Server{Addr: "localhost:8080", Handler: &handler}
	server.ListenAndServe()

}

func rktHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hit rocket endpoint")

}
