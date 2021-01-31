package database

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const (
	CONNECTURI      = "mongodb://localhost:27017"
	DB              = "pi-hub"
	SENSORSSECURITY = "sensors/security"
)

func GetClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		o := options.Client().ApplyURI(CONNECTURI)
		client, err := mongo.Connect(context.TODO(), o)
		if err != nil {
			clientInstanceError = err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
