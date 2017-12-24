package main

import "gopkg.in/mgo.v2/bson"

// Tracker -
type Tracker struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`
}
