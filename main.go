package main

import (
	"github.com/IamNator/thealth/internal/models"
	"github.com/IamNator/thealth/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	defaultRouter := gin.Default()
	DB := models.SetupModels()

	r := router.New(defaultRouter, DB)
	r.AttachDB()
	defer r.CloseDB()

	r.PatientRouter()
	r.PhysicianRouter()
	r.LocalFacilityRouter()

	port := ":" + viper.GetString("PORT")
	r.Engine.Run(port)

}
