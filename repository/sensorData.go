package repository

import (
	"home-collect/database"
	"home-collect/domain"
)

type SensorDataRepository struct {
	*database.DB
}

type ISensorDataRepository interface {
	InsertSensorData(url string, data domain.SensorData) error
}

func (sd *SensorDataRepository) InsertSensorData(url string, data domain.SensorData) error {
	sess := sd.Session.Copy()
	defer sess.Close()

	err := sess.DB(sd.Database).C(url).Insert(&data)

	return err
}
