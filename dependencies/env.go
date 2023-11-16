package dependencies

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvPort() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv("PORT")
}
