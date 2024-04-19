package main

import (
	"contrack-api/models"
	"contrack-api/router"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvVariables()

	models.ConnectDatabase()

	app := router.Init()

	err := app.Run(":8080")

	if err != nil {
		panic("Failed to start the server!")
	}
}

func loadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env variables!")
	}
}
