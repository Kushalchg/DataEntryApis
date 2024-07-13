package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	fmt.Print("loaded successfully")
	if err != nil {
		log.Fatal("error while loading env variables")
	}
}
