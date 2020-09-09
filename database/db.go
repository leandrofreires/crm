package database

import (
	"context"
	"fmt"

	"github.com/leandrofreires/crm/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Db is connection with datasource
var Db *mongo.Database

//Connect Returne database connected
func Connect() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Env.GetDatabaseURL()))
	if err != nil {
		panic(err)
	}
	// Ping the primary
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return client
}
