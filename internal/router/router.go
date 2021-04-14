package router

import (
	"github.com/IamNator/thealth/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type router struct {
	Engine *gin.Engine
	DB     *gorm.DB
}

func New(e *gin.Engine, db *gorm.DB) router {
	return router{e, db}
}

//PatientRouter routes the different requests
//related to patients profile
func (r *router) PatientRouter() {
	r.Engine.GET("/patients", controllers.FindPatients)
	r.Engine.POST("/patients", controllers.CreatePatient)  //create
	r.Engine.GET("/patients/:id", controllers.FindPatient) //find by id

	r.Engine.PATCH("/patients/:id", controllers.UpdatePatient) //update by id
	r.Engine.DELETE("patients/:id", controllers.DeletePatient) //delete by id

}

//PhysicianRouter routes the different requests
//related to consult physicians profile
func (r *router) PhysicianRouter() {

	r.Engine.GET("/physicians", controllers.FindConsultPhyss)
	r.Engine.POST("/physicians", controllers.CreateConsultPhys)  //create
	r.Engine.GET("/physicians/:id", controllers.FindConsultPhys) //find by id

	r.Engine.PATCH("/physicians/:id", controllers.UpdateConsultPhys) //update by id
	r.Engine.DELETE("physicians/:id", controllers.DeleteConsultPhys) //delete by id

}



//PhysicianRouter routes the different requests
//related to consult physicians profile
func (r *router) LocalFacilityRouter() {

	r.Engine.GET("/physicians", controllers.FindLocalFacilities)
	r.Engine.POST("/physicians", controllers.CreateLocalFacility)  //create
	r.Engine.GET("/physicians/:id", controllers.FindLocalFacility) //find by id

	r.Engine.PATCH("/physicians/:id", controllers.UpdateLocalFacility) //update by id
	r.Engine.DELETE("physicians/:id", controllers.DeleteLocalFacility) //delete by id

}

func (r *router) AttachDB() {
	// db := models.SetupModels()

	// Provides db variable to controller
	r.Engine.Use(func(c *gin.Context) {
		c.Set("db", r.DB)
		c.Next()
	})

}

func (r *router) CloseDB() {
	r.DB.Close()
}
