package main

import (
	"fmt"
	"log"

	"github.com/afthaab/task-manager/pkg/db"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error: Error in loading .env file", err)
	}
}

func main() {
	db.ConnectToDatabase()
	fmt.Println("Main method invoked")
}
