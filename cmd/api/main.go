package main

import (
	"log"

	"github.com/afthaab/task-manager/pkg/di"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error: Error in loading .env file", err)
	}
}

func main() {
	server, err := di.InitializeApi()
	if err != nil {
		log.Fatalln("Could not start the Server: ", err)
	} else {
		server.Start()
	}
}
