package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PutPathParameter func NOTE PutPathParameter func
func PutPathParameter(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{"message": username})
}