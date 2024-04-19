package controllers

import (
	"contrack-api/dto"
	"contrack-api/models"
	"crypto/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func FindProjects(c *gin.Context) {
	var projects []models.Project
	models.DB.Find(&projects)

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func FindProject(c *gin.Context) {
	var project models.Project

	sub, _ := c.Get("sub")

	if sub == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized!"})
		return
	}

	if err := models.DB.Where("id = ?", sub).First(&project).Error; err != nil {
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

	// Create Project key and passkey
	key := generateKey(true, "")

	if key == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate project key!"})
		return
	}

	passKey := generatePassKey()

	if passKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate project pass key!"})
		return
	}

	// Hash passkey

	hash, err := bcrypt.GenerateFromPassword([]byte(passKey), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash pass key!"})
		return
	}

	project := models.Project{
		AgreedDeliveryDate: input.AgreedDeliveryDate,
		AgreedPrice:        input.AgreedPrice,
		OverallProgress:    input.OverallProgress,
		Key:                key,
		PassKey:            string(hash),
	}

	models.DB.Create(&project)

	c.JSON(http.StatusOK, gin.H{"data": dto.CreateProjectOutput{
		Key:     key,
		PassKey: passKey,
	}})
}

func generateKey(keyExists bool, prevString string) string {
	if !keyExists {
		return prevString
	}

	var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var length = 6
	b := make([]byte, length)
	_, err := rand.Read(b)

	if err != nil {
		return ""
	}

	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%len(chars)]
	}

	if err := models.DB.Where("key = ?", string(b)).First(&models.Project{}).Error; err != nil {
		return generateKey(false, string(b))
	} else {
		return generateKey(true, string(b))
	}
}

func generatePassKey() string {
	var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var length = 6
	b := make([]byte, length)
	_, err := rand.Read(b)

	if err != nil {
		return ""
	}

	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%len(chars)]
	}

	return string(b)
}
