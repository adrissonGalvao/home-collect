package database

import (
	mgo "gopkg.in/mgo.v2"
)

type DB struct {
	Session  *mgo.Session
	Database string
}
