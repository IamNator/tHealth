package controllers

import (
	"net/http"

	"github.com/IamNator/thealth/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//FindLocalProviders Gets all Local Providers information
//
// GET /local_facilities
func FindLocalProviders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var localProvider []models.LocalProvider
	db.Find(&localProvider)

	c.JSON(http.StatusOK, gin.H{"data": localProvider})
}

//CreateLocalProvider creates new Local Provider rec
//
// POST /local_providers
func CreateLocalProvider(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validate input
	var input models.LocalProvider
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Create physicians records
	var localProvider models.LocalProvider
	db.Model(&localProvider).Create(input)

	c.JSON(http.StatusCreated, gin.H{"data": localProvider})
}

//FindLocalProvider finds a single Local Providers profile
//
// GET /local_providers/:id
func FindLocalProvider(c *gin.Context) {
	//get db object
	db := c.MustGet("db").(*gorm.DB)

	//find local Providers profile
	var localProvider models.LocalProvider
	if err := db.Where("id = ?", c.Param("id")).First(&localProvider).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	//return a response
	c.JSON(http.StatusOK, gin.H{"data": localProvider})
}

//UpdateLocalProvider updates information of Local Providers
//
//PATCH /local_providers/:id
func UpdateLocalProvider(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get model if exists
	var localProvider models.LocalProvider
	if err := db.Where("id = ?", c.Param("id")).First(&localProvider).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	//validate input
	if err := c.ShouldBindJSON(&localProvider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//update records
	db.Model(&localProvider).Update(localProvider)

	c.JSON(http.StatusOK, gin.H{"data": localProvider})
}

//DeleteLocalProvider deletes a Local Provider
//
// DELETE /local_providers/:id
func DeleteLocalProvider(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get localProvider if exist
	var localProvider models.LocalProvider
	if err := db.Where("id = ? ", c.Param("id")).First(&localProvider).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	//delete localFacility record
	db.Delete(localProvider)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
