package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type JSONRespose struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type UserData struct {
	Status bool   `json:"status"`
	Id     string `json:"id"`
}

func Status(w http.ResponseWriter, _ *http.Request) {
	response := JSONRespose{
		Message: "Hellow its first",
		Status:  true,
	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func Root(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Welcome Back... with go...lang")
}

func AddData(w http.ResponseWriter, r *http.Request) {

	// get url data /{id}/
	reqData := mux.Vars(r)
	id := reqData["id"]

	response := UserData{
		Id:     id,
		Status: true,
	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(response)
}
