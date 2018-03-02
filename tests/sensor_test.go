package test

import (
	"home-collect/domain"
	"home-collect/service"
	mock "home-collect/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestVerifyIntegrityUrlSensor(t *testing.T) {
	sensorRepositoryMock := new(mock.ISensorRepository)

	sensorRepositoryMock.On("FindUrlSensor", "teste").Return(true, nil)

	sensorService := service.SensorService{sensorRepositoryMock}

	resultTest := sensorService.VerifyIntegrityUrlSensor("teste")

	assert.Equal(t, true, resultTest)
}

func TestGeneratingUrlSensors(t *testing.T) {
	sensorRepositoryMock := new(mock.ISensorRepository)

	var sensors []domain.Sensor

	sensor := domain.Sensor{}
	sensor.ID = bson.NewObjectId()
	sensor.Name = "teste"
	sensor.Token = "teste"
	sensor.Url = "teste"
	sensor.User = sensor.ID

	sensors = append(sensors, sensor)

	sensorRepositoryMock.On("FindAllSensor").Return(sensors, nil)

	sensorService := service.SensorService{sensorRepositoryMock}
	resultTest, _ := sensorService.GenerateUrlsSensor()

	assert.Equal(t, "/"+sensor.Url, resultTest[0])

}
