package main

import (
	"goAPIServer/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
