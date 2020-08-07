package handler

import (
	"github.com/gin-gonic/gin"
	"restfull/routers"
)

// Router API
func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/", routers.GetHelloWorld)
	router.POST("/", routers.PostAPI)
	router.PUT("/path/:username", routers.PutPathParameter)
	router.DELETE("/query", routers.DeleteQuery)

	return router
}
