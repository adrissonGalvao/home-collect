package servicecontainer

import (
	"home-collect/database"
	"home-collect/repository"
	"home-collect/service"
	"log"
	"sync"

	mgo "gopkg.in/mgo.v2"
)

type IServiceContainer interface {
	InjectUserService() service.IUserService
	InjectSensorService() service.ISersoService
	InjectSensorDataService() service.ISersoDataService
}

type kernel struct{}

const (
	DBSERVER = "mongodb://homecollect:homecollect@ds245548.mlab.com:45548/homecollect"
	DBNAME   = "homecollect"
)

func (k *kernel) InjectUserService() service.IUserService {
	session, err := mgo.Dial(DBSERVER)
	if err != nil {
		log.Fatal(err)
	}
	dbMongo := &database.DB{}
	dbMongo.Session = session
	dbMongo.Database = DBNAME

	userRepository := &repository.UserRepository{dbMongo}
	userService := &service.UserService{userRepository}

	return userService
}

func (k *kernel) InjectSensorService() service.ISersoService {
	session, err := mgo.Dial(DBSERVER)
	if err != nil {
		log.Fatal(err)
	}
	dbMongo := &database.DB{}
	dbMongo.Session = session
	dbMongo.Database = DBNAME

	sensorRepository := &repository.SensorRepository{dbMongo}
	sensorService := &service.SersoService{sensorRepository}

	return sensorService
}
func (k *kernel) InjectSensorDataService() service.ISersoDataService {
	session, err := mgo.Dial(DBSERVER)
	if err != nil {
		log.Fatal(err)
	}
	dbMongo := &database.DB{}
	dbMongo.Session = session
	dbMongo.Database = DBNAME

	sensorDataRepository := &repository.SensorDataRepository{dbMongo}
	sensorRepository := &repository.SensorRepository{dbMongo}
	sensorDataService := &service.SersoDataService{sensorDataRepository, sensorRepository}

	return sensorDataService
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
