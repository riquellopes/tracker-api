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
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Data      string        `bson:"data" json:"data"`
	Opened    bool          `bson:"opened" json:"opened"`
	CreateAt  time.Time     `bson:"create_at" json:"create_at"`
	UpdatedAt time.Time     `bson:"update_at" json:"update_at"`
}

// Add -
func (m *Tracker) Add() error {
	m.CreateAt = time.Now()
	m.ID = bson.NewObjectId()
	m.Opened = false

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
