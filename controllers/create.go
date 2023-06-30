package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../views"

	model "../models"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// take some data
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			// save it
			fmt.Println(data, "DATA")
			if err := model.CreateTodo(data.Name, data.ToDo); err != nil {
				println(err, "Error")
				w.Write([]byte("Some Error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
