package controllers

import (
	"net/http"

	"github.com/IamNator/thealth/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//FindPatients Gets all patients
//
// GET /patients
func FindPatients(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var patients []models.PatientProfiles
	db.Find(&patients)

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

//CreatePatient creates new patient
//
// POST /patients
func CreatePatient(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validate input
	var input models.CreatePatientProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Create Book
	patient := models.PatientProfiles{
		FullName:        input.FullName,
		Gender:          input.Gender,
		Contact:         input.Contact,
		ReligionCulture: input.ReligionCulture,
		Telephone:       input.Telephone,
		// Languages:       input.Languages,
		// Assessment:      input.Assessment,
		// History:         input.History,
		// Objective:       input.Objective,
		// AttachedFiles:   input.AttachedFiles,
		ResponseTime:    input.ResponseTime,
		TimeofCompalint: input.TimeofCompalint,
	}

	db.Create(patient)

	c.JSON(http.StatusCreated, gin.H{"data": patient})
}

//FindPatient finds a single patient profile
//
// GET /patients/:id
func FindPatient(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var patientProfile models.PatientProfiles
	if err := db.Where("id = ?", c.Param("id")).First(&patientProfile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patientProfile})
}

//UpdatePatient updates information of patient
//
//PATCH /patients/:id
func UpdatePatient(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get model if exists
	var patientProfile models.PatientProfiles
	if err := db.Where("id = ?", c.Param("id")).First(&patientProfile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	//validate input
	var input models.UpdatePatientProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//update records
	db.Model(&input).Update(input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

//DeletePatient deletes a patient
//
// DELETE /patients/:id
func DeletePatient(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get patient if exist
	var patient models.PatientProfiles
	if err := db.Where("id = ? ", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}

	//delete patient record
	db.Delete(patient)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
