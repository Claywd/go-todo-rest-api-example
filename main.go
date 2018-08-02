package main

import (
	"log"

	"github.com/endofcake/go-todo-rest-api-example/app"
	"github.com/endofcake/go-todo-rest-api-example/config"
)

func main() {
	log.Print("Starting execution...")
	log.Print("Getting config...")
	config := config.GetConfig()

	log.Print("Initialising app...")
	app := &app.App{}
	app.Initialize(config)

	log.Print("Starting the server...")
	app.Run(":3000")
}
