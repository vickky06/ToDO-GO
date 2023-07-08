package controller

import (
	"encoding/json"
	"net/http"

	model "../models"
)

func readByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, error := model.ReadByName(name)
			if error != nil {
				w.Write([]byte(error.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
