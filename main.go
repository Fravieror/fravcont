package main

import (
	"api-test/server"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := server.New()
	log.Fatal(http.ListenAndServe(os.Getenv("PORTAPI"), s.Router()))
}
