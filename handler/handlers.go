package handler

import (
	"github.com/gin-gonic/gin"
	"restfull/controllers"
)

// Router API
func Router() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", controllers.GetHelloWorld)
		v1.POST("/", controllers.PostAPI)
		v1.PUT("/", controllers.PutPathParameter)
		v1.DELETE("/", controllers.DeleteQuery)
	}
	return router
}
