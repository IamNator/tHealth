package router

import (
	"github.com/IamNator/thealth/internal/controllers"
	"github.com/IamNator/thealth/internal/models"
	"github.com/gin-gonic/gin"
)

type router struct {
	*gin.Engine
}

func New(e *gin.Engine) router {
	return router{e}
}

//PatientRouter routes the different requests
//related to patients profile
func (r *router) PatientRouter() {

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

}

//PhysicianRouter routes the different requests
//related to consult physicians profile
func (r *router) PhysicianRouter() {

	db := models.SetupModels()

	// Provides db variable to controller
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/physicians", controllers.FindConsultPhyss)
	r.POST("/physicians", controllers.CreateConsultPhys)  //create
	r.GET("/physicians/:id", controllers.FindConsultPhys) //find by id

	r.PATCH("/physicians/:id", controllers.UpdateConsultPhys) //update by id
	r.DELETE("physicians/:id", controllers.DeleteConsultPhys) //delete by id

}
