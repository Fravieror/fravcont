package infrastructure

import (
	"api-test/movement/application"
	"api-test/movement/domain"
	"encoding/json"
	"log"
	"net/http"
)

// Get only one movement for id
func Get(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if !ok {
		w.Write([]byte(`{"message": "param id not sended"}`))
		return
	}
	bussines := application.NewBussines(&MySQLRepository{})
	response := bussines.Get(id[0])
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Error cast struct to json"))
	}
	w.Write(jsonResponse)
}

// Save movement
func Post(w http.ResponseWriter, r *http.Request) {
	var movement domain.Movement
	bussines := application.NewBussines(&MySQLRepository{})
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&movement)
	log.Println(r.Body)
	log.Println(movement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := bussines.Save(&movement)
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		w.Write([]byte("Error cast struct to json"))
	}
	w.Write(jsonResponse)
}

// Get all movements for client
func GetList(w http.ResponseWriter, r *http.Request) {
	client, ok := r.URL.Query()["client"]
	if !ok {
		w.Write([]byte(`{"message": "param client not sended"}`))
		return
	}
	bussines := application.NewBussines(&MySQLRepository{})
	response := bussines.GetList(client[0])
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Error cast struct to json"))
	}
	w.Write(jsonResponse)
}

// Test function
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Pong"}`))
}
