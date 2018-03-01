package repository

import (
	"home-collect/database"
	"home-collect/domain"

	"gopkg.in/mgo.v2/bson"
)

const (
	DBSERVER = "mongodb://homecollect:homecollect@ds245548.mlab.com:45548/homecollect"
	DBNAME   = "homecollect"
)

type IUserRepository interface {
	InsertUser(user domain.User) error
	FindAllUser() ([]domain.User, error)
	FindByIdUser(id string) (domain.User, error)
	DeleteUser(user domain.User) error
	UpdateUser(user domain.User) error
}

type UserRepository struct {
	*database.DB
}

func (u *UserRepository) InsertUser(user domain.User) error {
	sess := u.Session.Copy()
	defer sess.Close()

	err := sess.DB(u.Database).C("user").Insert(&user)

	return err
}

func (u *UserRepository) FindAllUser() ([]domain.User, error) {
	sess := u.Session.Copy()
	defer sess.Close()
	var users []domain.User

	err := sess.DB(u.Database).C("user").Find(bson.M{}).All(&users)
	return users, err

}
func (u *UserRepository) FindByIdUser(id string) (domain.User, error) {
	sess := u.Session.Copy()
	defer sess.Clone()
	var user domain.User

	err := sess.DB(u.Database).C("user").FindId(bson.ObjectIdHex(id)).One(&user)

	return user, err
}

func (u *UserRepository) DeleteUser(user domain.User) error {
	sess := u.Session.Copy()
	defer sess.Close()
	err := sess.DB(u.Database).C("user").Remove(&user)

	return err
}

func (u *UserRepository) UpdateUser(user domain.User) error {
	sess := u.Session.Copy()
	defer sess.Close()

	err := sess.DB(u.Database).C("user").UpdateId(user.ID, &user)

	return err
}
