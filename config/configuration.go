package config

import (
	"github.com/joho/godotenv"
	"os"
)

type config struct {
	Port string
}

var Config config

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		println("Error when loading .env: " + err.Error())
	}
	Config.Port = os.Getenv("PORT")
}
