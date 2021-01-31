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
	Door struct {
		ID     primitive.ObjectID `bson:"_id" json:"_id"`
		Type   string             `json:"type" validate:"required,oneof=door.lock door.open"`
		Name   string             `json:"name" validate:"required"`
		State  DoorState          `json:"state,omitempty"`
		Online bool               `json:"online,omitempty"`
		MAC    string             `json:"mac,omitempty" validate:"omitempty,mac"`
	}

	DoorState struct {
		Open        bool      `bson:"open"`
		Unlocked    bool      `bson:"unlocked"`
		LastUpdated time.Time `bson:"last_updated"`
	}
)

func AddDoorSensorDoc(door Door) (string, error) {
	client, err := database.GetClient()
	if err != nil {
		return "", err
	}

	collection := client.Database(database.DB).Collection(database.SENSORSSECURITY)
	r, err := collection.InsertOne(context.TODO(), &Door{
		ID:   primitive.NewObjectID(),
		Type: door.Type,
		Name: door.Name,
		State: DoorState{
			Open:        false,
			Unlocked:    false,
			LastUpdated: time.Time{},
		},
		Online: false,
		MAC:    door.MAC,
	})
	if err != nil {
		return "", err
	}

	if oid, ok := r.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", nil
}

func GetDoorSensorDoc(id primitive.ObjectID) (Door, error) {
	result := Door{}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	client, err := database.GetClient()
	if err != nil {
		return result, err
	}

	collection := client.Database(database.DB).Collection(database.SENSORSSECURITY)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetAllDoorSensorsDocs() ([]Door, error) {
	var doors []Door

	filter := bson.D{primitive.E{Key: "type", Value: "security.sensor.door"}}
	client, err := database.GetClient()
	if err != nil {
		return doors, err
	}

	collection := client.Database(database.DB).Collection(database.SENSORSSECURITY)

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return nil, findError
	}

	for cur.Next(context.TODO()) {
		d := Door{}
		err := cur.Decode(&d)
		if err != nil {
			return doors, err
		}
		doors = append(doors, d)
	}

	cur.Close(context.TODO())
	if len(doors) == 0 {
		return doors, mongo.ErrNoDocuments
	}
	return doors, nil
}
