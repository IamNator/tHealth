package main

import (
	"github.com/IamNator/thealth/internal/models"
	"github.com/IamNator/thealth/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	defaultRouter := gin.Default() //default gin router
	DB := models.SetupModels()     //open up databse connection and run migration db does not exist

	r := router.New(defaultRouter, DB)
	r.AttachDB()      //Attach database connection
	defer r.CloseDB() //Close DB connection on apps termination

	r.Cors() //Apply cors

	r.PatientRouter()
	r.PhysicianRouter()
	r.LocalFacilityRouter()
	r.LocalProviderRouter()

	port := ":" + viper.GetString("PORT") //Get PORT addr from environ
	r.Engine.Run(port)                    //Listen and serve on PORT

}
