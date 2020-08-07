package router

import (
	// "github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)

// Router API
func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/", GetHelloWorld)
	router.POST("/", PostAPI)
	router.PUT("/path/:username", PutPathParameter)
	router.DELETE("/query", DeleteQuery)

	return router
}
