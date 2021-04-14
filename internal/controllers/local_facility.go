package controllers

import (
	"net/http"

	"github.com/IamNator/thealth/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//FindLocalFacilities Gets all Local Facilities information
//
// GET /local_facilities
func FindLocalFacilities(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var localFacility []models.LocalFacility
	db.Find(&localFacility)

	c.JSON(http.StatusOK, gin.H{"data": localFacility})
}

//CreateLocalFacilities creates new LocalFacilities
//
// POST /local_facilities
func CreateLocalFacility(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validate input
	var input models.LocalFacility
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Create physicians records
	var localFacility models.LocalFacility
	db.Model(&localFacility).Create(input)

	c.JSON(http.StatusCreated, gin.H{"data": localFacility})
}

//FindLocalFacility finds a single Local Facility profile
//
// GET /local_facilities/:id
func FindLocalFacility(c *gin.Context) {
	//get db object
	db := c.MustGet("db").(*gorm.DB)

	//find local Facility profile
	var localFacility models.LocalFacility
	if err := db.Where("id = ?", c.Param("id")).First(&localFacility).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	//return a response
	c.JSON(http.StatusOK, gin.H{"data": localFacility})
}

//UpdateLocalFacilities updates information of Local Facility
//
//PATCH /local_facilities/:id
func UpdateLocalFacility(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get model if exists
	var localFacility models.LocalFacility
	if err := db.Where("id = ?", c.Param("id")).First(&localFacility).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	//validate input
	if err := c.ShouldBindJSON(&localFacility); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//update records
	db.Model(&localFacility).Update(localFacility)

	c.JSON(http.StatusOK, gin.H{"data": localFacility})
}

//DeleteLocalFacility deletes a Local Facility
//
// DELETE /local_facilities/:id
func DeleteLocalFacility(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get localFacility if exist
	var localFacility models.LocalFacility
	if err := db.Where("id = ? ", c.Param("id")).First(&localFacility).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	//delete localFacility record
	db.Delete(localFacility)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
