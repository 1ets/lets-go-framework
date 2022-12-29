package lets

import (
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
)

// Loading .env environment variable into memory.
func loadEnv() error {
	golog.Info("Loading: .env ...")
	err := godotenv.Load()

	if err != nil {
		golog.Fatal("Error loading .env file")
	}

	return err
}
