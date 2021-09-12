package main

import (
	"log"
	"rsoi-2021-lab1-ci-cd-Kotyarich/server"
)

func main() {
	app := server.NewApp()

	if err := app.Run(":8080"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
