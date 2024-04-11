package controllers

import (
	"contrack-api/dto"
	"contrack-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindProjects(c *gin.Context) {
	var projects []models.Project
	models.DB.Find(&projects)

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func FindProject(c *gin.Context) {
	var project models.Project

	if err := models.DB.Where("id = ?", c.Param("id")).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

func CreateProject(c *gin.Context) {
	var input dto.CreateProjectInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{
		AgreedDeliveryDate: input.AgreedDeliveryDate,
		AgreedPrice:        input.AgreedPrice,
		OverallProgress:    input.OverallProgress,
	}

	models.DB.Create(&project)

	c.JSON(http.StatusOK, gin.H{"data": project})
}