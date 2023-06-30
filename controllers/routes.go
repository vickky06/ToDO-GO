package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	// println("Register called")
	mux.HandleFunc(
		"/ping", Ping())

	mux.HandleFunc("/createTodo", create())
	return mux
}
