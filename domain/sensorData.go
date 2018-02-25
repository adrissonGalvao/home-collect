package domain

import (
	"gopkg.in/mgo.v2/bson"
)

type SensorData struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Data   string        `bson:"data" json:"data"`
	Sensor bson.ObjectId `bson:"user $ref:sensor"`
}
