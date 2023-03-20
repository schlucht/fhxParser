package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/schlucht/fhxreader/server/controllers"
)

func FhxRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/fhx/units", controllers.Unitnames())
	incomingRoutes.GET("/fhx/up")
}
