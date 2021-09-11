package main

import (
	"../../server"
	"log"
)

func main() {
	app := server.NewApp()

	if err := app.Run(":5000"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
