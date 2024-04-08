package controllers

import (
	"contrack-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindProjects(c *gin.Context) {
	var projects []models.Project
	models.DB.Find(&projects)

	c.JSON(http.StatusOK, gin.H{"data": projects})
}