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
func (m *Tracker) Add(tracker Tracker) error {
	err := db.C("tracker").Insert(&tracker)
	return err
}

// Exists -
func (m *Tracker) Exists(id string) (bool, error) {
	var tracker Tracker
	error := db.C("tracker").FindId(bson.ObjectIdHex(id)).One(&tracker)

	if error != nil {
		return true, nil
	}

	return false, error
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

	log.Print("PAssando2")

	session.SetMode(mgo.Monotonic, true)
	// db = session.DB(database)
}
