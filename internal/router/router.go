package router

import (
	"github.com/IamNator/thealth/internal/controllers"
	"github.com/IamNator/thealth/internal/models"
	"github.com/gin-gonic/gin"
)

//PatientRouter routes the different requests
func PatientRouter(r *gin.Engine) *gin.Engine {

	db := models.SetupModels()

	// Provides db variable to controller
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/patients", controllers.FindPatients)
	r.POST("/patients", controllers.CreatePatient)  //create
	r.GET("/patients/:id", controllers.FindPatient) //find by id

	r.PATCH("/patients/:id", controllers.UpdatePatient) //update by id
	r.DELETE("patients/:id", controllers.DeletePatient) //delete by id

	return r
}
