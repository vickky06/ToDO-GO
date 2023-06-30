package controller

import (
	"encoding/json"
	"net/http"

	views "../views"
)

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// println("PING CALLED")
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
