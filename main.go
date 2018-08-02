package main

import (
	"log"

	"github.com/endofcake/go-todo-rest-api-example/app"
	"github.com/endofcake/go-todo-rest-api-example/config"
)

func main() {
	log.Print("Starting execution...")
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
