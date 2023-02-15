package main

import (
	"fmt"
	"log"

	"github.com/ariwiraa/notes-app/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("read .env from file")

	config.InitDb()
}
