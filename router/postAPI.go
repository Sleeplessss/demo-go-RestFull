package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
import ( "restfull/models" )
// PostAPI func
func PostAPI(c *gin.Context) {
	u := new(models.User)
	c.BindJSON(&u)
	c.JSON(http.StatusOK, gin.H{"user": u})
}