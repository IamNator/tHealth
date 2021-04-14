package main

import (
	"github.com/IamNator/thealth/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	rDefault := gin.Default()
	
	r := router.New(rDefault)
	r.PatientRouter()
	r.PhysicianRouter()

	port := ":" + viper.GetString("PORT")
	r.Run(port)
}
