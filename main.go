package main

import (
	"goAPIServer/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.InitRoutes()

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
