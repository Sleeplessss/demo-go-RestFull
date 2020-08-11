package controllers

import (
	"fmt"
	// "io/ioutil"
	"context"
	"log"
	"net/http"
	"restfull/configs"
	"restfull/models"

	// "reflect"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// PutPathParameter func NOTE PutPathParameter func
func PutPathParameter(c *gin.Context) {
	// connect to db
	configs.ConnectionMongoDB()

	body := new(models.User)
	c.BindJSON(&body)

	var setElements bson.D
	if body.Name != "" {
		setElements = append(setElements, bson.E{"name", body.Name})
	}
	if body.Surname != "" {
		setElements = append(setElements, bson.E{"surname", body.Surname})
	}

	setMap := bson.D{
		{"$set", setElements},
	}
	fmt.Printf("update: %+v\n", setElements)

	collection := configs.ClientMongoDB.Database("demoGolang").Collection("user")
	// err := collection.FindOne(context.TODO(), bson.M{"_id": body.ID}).Decode(&result)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": body.ID}, setMap)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Updated %v Documents!\n", updateResult.ModifiedCount)
	// fmt.Printf("Updated %v\n", updateResult.UpsertedID)

	var result models.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": body.ID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	// Disconnect db
	defer configs.DisconnectMongoDB()
	c.JSON(http.StatusOK, gin.H{"result": result})
}
