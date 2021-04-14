package controllers

import (
	"net/http"

	"github.com/IamNator/thealth/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//FindConsultPhyss Gets all physicians information
//
// GET /physicians
func FindConsultPhyss(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var physician []models.ConsultPhys
	db.Find(&physician)

	c.JSON(http.StatusOK, gin.H{"data": physician})
}

//CreateConsultPhys creates new ConsultPhys
//
// POST /physicians
func CreateConsultPhys(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validate input
	var input models.CreateConsultPhys
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Create physicians records
	var ConsultPhys models.ConsultPhys
	db.Model(&ConsultPhys).Create(input)

	c.JSON(http.StatusCreated, gin.H{"data": ConsultPhys})
}

//FindConsultPhys finds a single ConsultPhys profile
//
// GET /physicians/:id
func FindConsultPhys(c *gin.Context) {
	//get db object
	db := c.MustGet("db").(*gorm.DB)

	//finc consultants profile
	var consultPhys models.ConsultPhys
	if err := db.Where("id = ?", c.Param("id")).First(&consultPhys).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	//return a response
	c.JSON(http.StatusOK, gin.H{"data": consultPhys})
}

//UpdateConsultPhys updates information of ConsultPhys
//
//PATCH /physicians/:id
func UpdateConsultPhys(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get model if exists
	var consultPhys models.ConsultPhys
	if err := db.Where("id = ?", c.Param("id")).First(&consultPhys).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	//validate input
	if err := c.ShouldBindJSON(&consultPhys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//update records
	db.Model(&consultPhys).Update(consultPhys)

	c.JSON(http.StatusOK, gin.H{"data": consultPhys})
}

//DeleteConsultPhys deletes a physician
//
// DELETE /physicians/:id
func DeleteConsultPhys(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get physician if exist
	var consultPhys models.ConsultPhys
	if err := db.Where("id = ? ", c.Param("id")).First(&consultPhys).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	//delete physician record
	db.Delete(consultPhys)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
