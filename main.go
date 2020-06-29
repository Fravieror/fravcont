package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/save", save)
	http.HandleFunc("/get", getList)
	http.HandleFunc("/delete", delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func save(w http.ResponseWriter, r *http.Request) {
	body, _ := r.GetBody()
	bussines := bussiness{}
	w.Write([]byte(`{"message": "hello world"}`))
}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}
