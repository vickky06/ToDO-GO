package main

import (
	"log"
	"net/http"

	controller "./controllers"
	model "./models"
)

func main() {
	mux := controller.Register()
	db, err := model.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.ListenAndServe(":3000", mux)
}
