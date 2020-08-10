package controllers

import (
	"context"
	"log"
	"net/http"
	"restfull/configs"
	"restfull/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetHelloWorld func NOTE GetHelloWorld func
func GetHelloWorld(c *gin.Context) {

	// connect db
	configs.ConnectionMongoDB()

	findOptions := options.Find()
	var results []*models.User
	collection := configs.ClientMongoDB.Database("demoGolang").Collection("user")
	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
		// defer cur.Close(context.TODO())
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	defer configs.DisconnectMongoDB()
	c.JSON(http.StatusOK, gin.H{"result": results})
}
