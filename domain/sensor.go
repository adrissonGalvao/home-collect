package domain

import (
	"gopkg.in/mgo.v2/bson"
)

type Sensor struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Token string        `bson:"token" json:"token"`
	Url   string        `bson:"url" json:"url"`
	User  bson.ObjectId `bson:"user $ref:user" json:"user"`
}
