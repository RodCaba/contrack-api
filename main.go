package main

import (
	"contrack-api/controllers"
	"contrack-api/models"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/projects", controllers.FindProjects)

	router.Run("localhost:8080")
}
