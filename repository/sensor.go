package repository

import (
	"home-collect/domain"
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type SensorRepository struct {
	database string
	session  *mgo.Session
}

func (s *SensorRepository) Connect() {
	session, err := mgo.Dial(DBSERVER)
	s.database = DBNAME
	if err != nil {
		log.Fatal(err)
	}
	s.session = session
}

func (s *SensorRepository) InsertSensor(sensor domain.Sensor) error {
	sess := s.session.Copy()
	defer sess.Close()

	err := sess.DB(s.database).C("sensor").Insert(&sensor)

	return err
}

func (s *SensorRepository) FindAllSensor() ([]domain.Sensor, error) {
	sess := s.session.Copy()
	defer sess.Close()
	var sensors []domain.Sensor

	err := sess.DB(s.database).C("sensor").Find(bson.M{}).All(&sensors)

	return sensors, err
}

func (s *SensorRepository) FindByIdSensor(id string) (domain.Sensor, error) {
	sess := s.session.Copy()
	defer sess.Close()

	var sensor domain.Sensor

	err := sess.DB(s.database).C("sensor").FindId(bson.ObjectIdHex(id)).One(&sensor)

	return sensor, err
}

func (s *SensorRepository) DeleteSensor(senso domain.Sensor) error {
	sess := s.session.Copy()
	defer sess.Close()

	err := sess.DB(s.database).C("sensor").Remove(&senso)

	return err
}
func (s *SensorRepository) UpdateSensor(sensor domain.Sensor) error {
	sess := s.session.Copy()
	defer sess.Close()

	err := sess.DB(s.database).C("sensor").UpdateId(sensor.ID, &sensor)

	return err
}

func (s *SensorRepository) FindUrlSensor(url string) (bool, error) {
	sess := s.session.Copy()
	defer sess.Close()
	count, err := sess.DB(s.database).C("sensor").Find(bson.M{"url": url}).Count()
	if count > 0 {
		return false, err
	}
	return true, err

}
