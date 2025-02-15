package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type enviroment struct {
	dir string
}

var env enviroment

func makeEnv() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
	}

	// Getting and using a value from .env
	env.dir = os.Getenv("DIRECTORY")

	fmt.Printf("These are the enviroment variables:\ndir | %s\n", env.dir)
}
