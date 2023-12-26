package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func main() {
	fmt.Println("Hello World")
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("empty port")
		return
	}
	apiConfig := Config{
		PORT: port,
	}

	serverEntry(&apiConfig)
}
