package controller

import (
	"encoding/json"
	"net/http"

	model "../models"
)

func readAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data, err := model.ReadAll()
			if err != nil {
				w.Write([]byte("Error occured while reading Data"))
				w.WriteHeader(http.StatusExpectationFailed)
				return
			}
			w.WriteHeader((http.StatusOK))
			json.NewEncoder(w).Encode(data)
		}
	}

}
