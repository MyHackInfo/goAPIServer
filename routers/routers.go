package routers

import (
	"goAPIServer/routers/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Root).Methods("GET")
	router.HandleFunc("/status", handlers.Status).Methods("GET")
	router.HandleFunc("/{id}/AddData", handlers.AddData).Methods("POST")

	return router
}
