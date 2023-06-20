package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pinged"))
	})
	http.ListenAndServe("localhost:3000", mux)
}

/***
package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method == http.MethodGet {
		// 	// structs.Response
		// }
		w.WriteHeader(http.StatusNotFound)
	})
	http.ListenAndServe("localhost:3000", mux)
}

/