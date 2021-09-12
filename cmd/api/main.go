package main

import (
	"log"
	"os"
	"rsoi-2021-lab1-ci-cd-Kotyarich/server"
)

func main() {
	port := os.Getenv("PORT")
	app := server.NewApp()

	if err := app.Run(":" + port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
