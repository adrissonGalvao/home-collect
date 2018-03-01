package mock

import (
	"home-collect/domain"

	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2/bson"
)

type ISensorRepository struct {
	mock.Mock
}

func (sm *ISensorRepository) InsertSensor(sensor domain.Sensor) error {
	args := sm.Called(sensor)

	return args.Error(0)
}

func (sm *ISensorRepository) FindAllSensor() ([]domain.Sensor, error) {
	args := sm.Called()
	var sensors []domain.Sensor
	sensor := domain.Sensor{}
	sensor.ID = bson.NewObjectId()
	sensor.Name = "teste"
	sensor.Token = "teste"
	sensor.Url = "teste"
	sensor.User = sensor.ID
	sensors[0] = sensor
	return sensors, args.Error(1)
}

func (sm *ISensorRepository) FindByIdSensor(id string) (domain.Sensor, error) {
	args := sm.Called()
	var sensor domain.Sensor
	return sensor, args.Error(1)
}

func (sm *ISensorRepository) DeleteSensor(sensor domain.Sensor) error {
	args := sm.Called(sensor)

	return args.Error(0)
}
func (sm *ISensorRepository) UpdateSensor(sensor domain.Sensor) error {
	args := sm.Called(sensor)

	return args.Error(0)
}

func (sm *ISensorRepository) FindUrlSensor(url string) (bool, error) {
	args := sm.Called(url)

	return args.Bool(0), args.Error(1)

}

func (sm *ISensorRepository) FindIdSensorByUrl(url string) (bson.ObjectId, error) {
	args := sm.Called(url)
	bson := bson.NewObjectId()
	return bson, args.Error(1)
}
