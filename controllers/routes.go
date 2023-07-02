package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	// println("Register called")
	mux.Handle(
		"/ping", Ping())

	mux.Handle("/createTodo", create())
	mux.Handle("/readAll", readAll())
	mux.Handle("/readByName", readByName())
	mux.Handle("/delete", delete())
	return mux
}
