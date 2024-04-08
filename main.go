package main

import (
	"contrack-api/models"
	"contrack-api/router"
)

type Project struct {
	ID int `json:"id"`
	AgreedDeliveryDate string `json:"agreedDeliveryDate"`
	LastEstimate float64 `json:"lastEstimate"`
	LastEstimateDate string `json:"lastEstimateDate"`
	AgreedPrice float64 `json:"agreedPrice"`
	OverallProgress float64 `json:"overallProgress"`
}

var projects = []Project{
	{ID: 1, AgreedDeliveryDate: "2021-01-01", LastEstimate: 1000, LastEstimateDate: "2021-01-01", AgreedPrice: 1000, OverallProgress: 0.5},
	{ID: 2, AgreedDeliveryDate: "2021-01-01", LastEstimate: 1000, LastEstimateDate: "2021-01-01", AgreedPrice: 1000, OverallProgress: 0.5},
}

func main() {

	models.ConnectDatabase()

	app := router.Init()

	err := app.Run(":8080")

	if err != nil {
		panic("Failed to start the server!")
	}
}
