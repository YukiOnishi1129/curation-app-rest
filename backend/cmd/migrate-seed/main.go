package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Start migrate seed")
	// ctx := context.Background()
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env.local file: %v", err)
		return
	}
	// db, err := database.Init()
	// if err != nil {
	// 	log.Fatalf("Failed to init db: %v", err)
	// 	return
	// }

	log.Printf("End migrate seed")

}
