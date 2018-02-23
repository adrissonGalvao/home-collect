package repository

import (
	"home-collect/domain"
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

const (
	DBSERVER = "mongodb://homecollect:homecollect@ds245548.mlab.com:45548/homecollect"
	DBNAME   = "homecollect"
)

type UserRepository struct {
	database string
	session  *mgo.Session
}

func (u *UserRepository) Connect() {
	session, err := mgo.Dial(DBSERVER)
	u.database = DBNAME
	if err != nil {
		log.Fatal(err)
	}
	u.session = session
}
func (u *UserRepository) InsertUser(user domain.User) error {
	sess := u.session.Copy()
	defer sess.Close()

	err := sess.DB(u.database).C("user").Insert(&user)

	return err
}

func (u *UserRepository) FindAllUser() ([]domain.User, error) {
	sess := u.session.Copy()
	defer sess.Close()
	var users []domain.User

	err := sess.DB(u.database).C("user").Find(bson.M{}).All(&users)
	return users, err

}
func (u *UserRepository) FindByIdUser(id string) (domain.User, error) {
	sess := u.session.Copy()
	defer sess.Clone()
	var user domain.User

	err := sess.DB(u.database).C("user").FindId(bson.ObjectIdHex(id)).One(&user)

	return user, err
}

func (u *UserRepository) DeleteUser(user domain.User) error {
	sess := u.session.Copy()
	defer sess.Close()
	err := sess.DB(u.database).C("user").Remove(&user)

	return err
}

func (u *UserRepository) UpdateUser(user domain.User) error {
	sess := u.session.Copy()
	defer sess.Close()

	err := sess.DB(u.database).C("user").UpdateId(user.ID, &user)

	return err
}
