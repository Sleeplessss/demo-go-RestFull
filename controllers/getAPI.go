package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetHelloWorld func NOTE GetHelloWorld func
func GetHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello Go lang"})
}