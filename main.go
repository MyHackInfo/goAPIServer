package main

import (
	"log"
	"net/http"
	"web/routers"
)

func main() {

	router := routers.InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
