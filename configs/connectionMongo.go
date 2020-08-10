package configs

import (
	"log"
	"fmt"
	"context"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ClientMongoDB global variable
var ClientMongoDB *mongo.Client

// ConnectionMongoDB server
func ConnectionMongoDB() {

	// load enviroment to connect db
	client, errConnection := GetClient()
	if errConnection != nil {
		log.Fatal(errConnection)
	}

	// check connection
	// errCheckConnection := CheckConnection(client)
	// if errCheckConnection != nil {
	// 	log.Fatal(errCheckConnection)
	// }

	//set global value
	ClientMongoDB = client

}

// GetClient for connect db
func GetClient() (*mongo.Client, error)  {

	// config client for connect to db
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", os.Getenv("MONGO_HOST"))).
  SetAuth(options.Credential{
    AuthSource: fmt.Sprintf("%s", os.Getenv("MONGO_DBNAME")), Username: fmt.Sprintf("%s", os.Getenv("MONGO_USERNAME")), Password: fmt.Sprintf("%s", os.Getenv("MONGO_PASSWORD")),
	})

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	
	if err != nil{
		fmt.Println("Error GetClient")
		log.Fatal(err)
	}

	errConnection := client.Ping(context.TODO(), nil)
	if errConnection != nil {
		fmt.Println("Error CheckConnection")
		log.Fatal(errConnection)
	}
	fmt.Println("Connected to MongoDB!")

	return client, nil
}

// CheckConnection ..
// func CheckConnection(client *mongo.Client) error  {
// 	errConnection := client.Ping(context.TODO(), nil)
// 	if errConnection != nil {
// 		fmt.Println("Error CheckConnection")
// 		log.Fatal(errConnection)
// 	}
// 	fmt.Println("Connected to MongoDB!")
// 	return nil
// }

// DisconnectMongoDB ..
func DisconnectMongoDB()  {
	err := ClientMongoDB.Disconnect(context.TODO())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB is Closed.")
}
