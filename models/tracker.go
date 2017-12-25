package models

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

// Tracker -
type Tracker struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`
}

// Add -
func (m *Tracker) Add() error {
	err := db.C("tracker").Insert(&m)
	return err
}

// Exists -
func (m *Tracker) Exists(id string) (bool, error) {
	count, error := db.C("tracker").Find(bson.M{"_id": bson.ObjectIdHex(id)}).Count()

	if error != nil || count > 0 {
		return true, nil
	}
	return false, nil
}

// Connection -
func Connection(username, password, server, database string) {
	info := &mgo.DialInfo{
		Addrs:    []string{server},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatal(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db = session.DB(database)
}
