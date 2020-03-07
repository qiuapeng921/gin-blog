package system

import (
	"github.com/joho/godotenv"
	"log"
)

func SetUp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}
