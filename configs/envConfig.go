package configs

import (
	"fmt"
	"github.com/joho/godotenv"
)

func InitEnvironment() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	
}
