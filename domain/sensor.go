package domain

import (
	"gopkg.in/mgo.v2/bson"
)

type Sensor struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	data  int64         `bson:"data" json:"data"`
	name  string        `bson:"name" json:"name"`
	token string        `bson:"token" json:"token"`
	url   string        `bson:"url" json:"url"`
	user  bson.ObjectId `bson:"user $ref:user" json:"user"`
}
