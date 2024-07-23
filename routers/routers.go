package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", root).Methods("GET")
	return router
}

func root(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Welcome Back... with go...lang")
}
