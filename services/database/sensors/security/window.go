package security

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"pi-hub/services/database"
	"time"
)

type (
	Window struct {
		ID     primitive.ObjectID `bson:"_id"`
		Type   string             `bson:"type"`
		Name   string             `bson:"name"`
		State  WindowState        `bson:"state"`
		Online bool               `bson:"online"`
	}

	WindowState struct {
		Open        bool      `bson:"open"`
		LastUpdated time.Time `bson:"last_updated"`
	}
)

func AddWindowSensor(door Window) error {
	client, err := database.GetClient()
	if err != nil {
		return err
	}
	collection := client.Database(database.DB).Collection(database.SENSORSSECURITY)
	_, err = collection.InsertOne(context.TODO(), door)

	if err != nil {
		return err
	}
	return nil
}

func GetWindowSensor(id primitive.ObjectID) ([]Window, error) {
	var windows []Window

	filter := bson.D{primitive.E{Key: "type", Value: "security.sensor.window"}}
	client, err := database.GetClient()
	if err != nil {
		return windows, err
	}

	collection := client.Database(database.DB).Collection(database.SENSORSSECURITY)

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return nil, findError
	}

	for cur.Next(context.TODO()) {
		w := Window{}
		err := cur.Decode(&w)
		if err != nil {
			return windows, err
		}
		windows = append(windows, w)
	}

	cur.Close(context.TODO())
	if len(windows) == 0 {
		return windows, mongo.ErrNoDocuments
	}
	return windows, nil
}
