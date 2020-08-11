package controllers

import (
	"context"
	"log"
	"net/http"
	"restfull/configs"
	"restfull/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteQuery func NOTE delete query func
func DeleteQuery(c *gin.Context) {
	// connect to db
	configs.ConnectionMongoDB()

	body := new(models.User)
	c.BindJSON(&body)
	collection := configs.ClientMongoDB.Database("demoGolang").Collection("user")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": body.ID})
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect db
	defer configs.DisconnectMongoDB()
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
