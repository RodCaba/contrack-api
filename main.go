package main

import (
	"contrack-api/models"
	"contrack-api/router"
)

func main() {

	models.ConnectDatabase()

	app := router.Init()

	err := app.Run(":8080")

	if err != nil {
		panic("Failed to start the server!")
	}
}
