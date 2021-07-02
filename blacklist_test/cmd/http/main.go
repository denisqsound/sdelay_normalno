package main

import (
	"fmt"
	"github.com/denisqsound/sdelay_normalno/blacklist_test/api"
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
	auth := api.GetAuth()
	fmt.Println(auth)

}
