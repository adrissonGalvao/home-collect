package test

import (
	"home-collect/domain"
	"home-collect/service"
	mock "home-collect/tests/mocks"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func TestVerifyIntegrityUrlSensor(t *testing.T) {
	sensorRepositoryMock := new(mock.ISensorRepository)

	sensorRepositoryMock.On("FindUrlSensor", "teste").Return(true, nil)

	sensorService := service.SensorService{sensorRepositoryMock}

	Convey("Sending URL: Teste", t, func() {
		Convey("Checking integruty of url sent", func() {
			resultTest := sensorService.VerifyIntegrityUrlSensor("teste")
			Convey("Validated url", func() {
				So(resultTest, ShouldEqual, true)
			})
		})
	})

}

func TestGenerateUrlsSensor(t *testing.T) {
	sensorRepositoryMock := new(mock.ISensorRepository)
	Convey("Creating sensor for test", t, func() {
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
		Convey("Generating Url for routes", func() {
			resultTest, _ := sensorService.GenerateUrlsSensor()
			Convey("verifying that the URLs were created correctly", func() {
				Convey("Url gererated correctly", func() {
					So("/"+sensors[0].Url, ShouldEqual, resultTest[0])
				})
			})
		})
	})
}
