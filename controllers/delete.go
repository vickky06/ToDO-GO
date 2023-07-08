package controller

import (
	"encoding/json"
	"net/http"

	model "../models"
)

func delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			name := r.URL.Path[1:] //path manupulation
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte("Error while Deletion"))
				return
			}
			json.NewEncoder(w).Encode(
				struct {
					Status string `json:status`
				}{`Item Deleted`})

		}
	}
}
