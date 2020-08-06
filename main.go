package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	"net/http"
	// "encoding/json"
)

// User struct NOTE user struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetHelloWorld func NOTE GetHelloWorld func
func GetHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

// Post func NOTE post func
func Post(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	c.JSON(http.StatusOK, gin.H{"user": u})
}

// PutPathParameter func NOTE PutPathParameter func
func PutPathParameter(c *gin.Context) {
	user := c.Param("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}

// DeleteQuery func NOTE delete query func
func DeleteQuery(c *gin.Context) {
	user := c.Query("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}

func main() {
	r := gin.Default()

	r.GET("/", GetHelloWorld)
	r.POST("/", Post)
	// r.GET("/", GetHelloWorld)
	// r.GET("/", GetHelloWorld)

	r.Run() // Run server
}
