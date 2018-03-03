package repository

import (
	"home-collect/database"
	"home-collect/domain"

	"gopkg.in/mgo.v2/bson"
)

type SensorDataRepository struct {
	*database.DB
}

type ISensorDataRepository interface {
	InsertSensorData(url string, data domain.SensorData) error
	FindAllSensorData(url string) ([]domain.SensorData, error)
	FindByIdSendorData(id string, url string) (domain.SensorData, error)
}

func (sd *SensorDataRepository) InsertSensorData(url string, data domain.SensorData) error {
	sess := sd.Session.Copy()
	defer sess.Close()

	err := sess.DB(sd.Database).C(url).Insert(&data)

	return err
}

func (sd *SensorDataRepository) FindAllSensorData(url string) ([]domain.SensorData, error) {
	sess := sd.Session.Copy()
	defer sess.Close()
	var sensorsData []domain.SensorData
	err := sess.DB(sd.Database).C(url).Find(bson.M{}).All(&sensorsData)

	return sensorsData, err
}

func (sd *SensorDataRepository) FindByIdSendorData(id string, url string) (domain.SensorData, error) {
	sess := sd.Session.Copy()
	defer sess.Close()

	var sensorData domain.SensorData

	err := sess.DB(sd.Database).C(url).FindId(bson.ObjectIdHex(url)).One(&sensorData)

	return sensorData, err
}
