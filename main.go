package main

import (
	"github.com/IamNator/thealth/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	rDefault := gin.Default()
	r := router.PatientRouter(rDefault)
	r.Run(":3000")
}
