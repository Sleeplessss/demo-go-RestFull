package controllers

import (
	"net/http"
	"fmt"
	"log"
	"context"
	"restfull/configs"
	"github.com/gin-gonic/gin"
	"restfull/models"
)
// PostAPI func
func PostAPI(c *gin.Context) {

	/*
		connect to db
	*/
	configs.ConnectionMongoDB()

	u := new(models.User)
	c.BindJSON(&u)
	// fmt.Println(u.Name)
	collection := configs.ClientMongoDB.Database("demoGolang").Collection("user")
	insertResult, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
			log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	defer configs.DisconnectMongoDB()
	c.JSON(http.StatusOK, gin.H{"user": u})
}