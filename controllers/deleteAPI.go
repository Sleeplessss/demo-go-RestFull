package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteQuery func NOTE delete query func
func DeleteQuery(c *gin.Context) {
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{"message": username})
}