package repository

import (
	"home-collect/database"
	"home-collect/domain"

	"gopkg.in/mgo.v2/bson"
)

type SensorRepository struct {
	*database.DB
}

type ISensorRepository interface {
	InsertSensor(sensor domain.Sensor) error
	FindAllSensor() ([]domain.Sensor, error)
	FindByIdSensor(id string) (domain.Sensor, error)
	DeleteSensor(sensor domain.Sensor) error
	UpdateSensor(sensor domain.Sensor) error
	FindIdSensorByUrl(url string) (bson.ObjectId, error)
	FindUrlSensor(url string) (bool, error)
}

func (s *SensorRepository) InsertSensor(sensor domain.Sensor) error {
	sess := s.Session.Copy()
	defer sess.Close()

	err := sess.DB(s.Database).C("sensor").Insert(&sensor)

	return err
}

func (s *SensorRepository) FindAllSensor() ([]domain.Sensor, error) {
	sess := s.Session.Copy()
	defer sess.Close()
	var sensors []domain.Sensor

	err := sess.DB(s.Database).C("sensor").Find(bson.M{}).All(&sensors)

	return sensors, err
}

func (s *SensorRepository) FindByIdSensor(id string) (domain.Sensor, error) {
	sess := s.Session.Copy()
	defer sess.Close()

	var sensor domain.Sensor

	err := sess.DB(s.Database).C("sensor").FindId(bson.ObjectIdHex(id)).One(&sensor)

	return sensor, err
}

func (s *SensorRepository) DeleteSensor(sensor domain.Sensor) error {
	sess := s.Session.Copy()
	defer sess.Close()

	err := sess.DB(s.Database).C("sensor").Remove(&sensor)

	return err
}
func (s *SensorRepository) UpdateSensor(sensor domain.Sensor) error {
	sess := s.Session.Copy()
	defer sess.Close()

	err := sess.DB(s.Database).C("sensor").UpdateId(sensor.ID, &sensor)

	return err
}

func (s *SensorRepository) FindUrlSensor(url string) (bool, error) {
	sess := s.Session.Copy()
	defer sess.Close()
	count, err := sess.DB(s.Database).C("sensor").Find(bson.M{"url": url}).Count()
	if count > 0 {
		return false, err
	}
	return true, err

}

func (s *SensorRepository) FindIdSensorByUrl(url string) (bson.ObjectId, error) {
	sess := s.Session.Copy()
	defer sess.Close()
	var sensor domain.Sensor

	err := sess.DB(s.Database).C("sensor").Find(bson.M{"url": url}).One(&sensor)

	return sensor.ID, err
}
