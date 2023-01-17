package loader

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// Loading .env environment variable into memory.
func Environment() {
	fmt.Println("Loading: .env ...")
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
