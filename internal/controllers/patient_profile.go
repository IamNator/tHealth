package controllers

import (
	"net/http"

	"github.com/IamNator/thealth/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//FindPatient finds a single patient profile
func FindPatient(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var patientProfile models.PatientProfile
	if err := db.Where("id = ?", c.Param("id")).First(&patientProfile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patientProfile})
}
