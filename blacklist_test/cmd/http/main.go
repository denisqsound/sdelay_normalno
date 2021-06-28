package main

import (
	"github.com/denisqsound/sdelay_normalno/blacklist_test/http"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	new(http.GetRequest)

}
