package handler

import (
	"github.com/gin-gonic/gin"
	"restfull/controllers"
)

// Router API
func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/", controllers.GetHelloWorld)
	router.POST("/", controllers.PostAPI)
	router.PUT("/path/:username", controllers.PutPathParameter)
	router.DELETE("/query", controllers.DeleteQuery)

	return router
}
