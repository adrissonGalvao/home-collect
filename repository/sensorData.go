package repository

import (
	"home-collect/domain"
	"log"

	mgo "gopkg.in/mgo.v2"
)

type SensorDataRepository struct {
	database string
	session  *mgo.Session
}

func (sd *SensorDataRepository) Connect() {
	session, err := mgo.Dial(DBSERVER)
	sd.database = DBNAME
	if err != nil {
		log.Fatal(err)
	}
	sd.session = session
}

func (sd *SensorDataRepository) InsertSensorData(url string, data domain.SensorData) error {
	sess := sd.session.Copy()
	defer sess.Close()

	err := sess.DB(sd.database).C(url).Insert(&data)

	return err
}
