package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	// fmt.Printf("Printing Error: %v\n", err)
	if err != nil {
		log.Fatal("Error loading .env files")
	}
}
