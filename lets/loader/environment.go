package loader

import (
	"log"

	"github.com/joho/godotenv"
)

// Loading .env environment variable into memory.
func Environment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
