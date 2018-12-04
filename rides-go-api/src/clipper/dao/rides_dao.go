package dao

import (
	"log"

	. "clipper/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RidesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "rides"
)

// Establish a connection to database
func (r *RidesDAO) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Find list of rides
func (r *RidesDAO) FindAll() ([]Ride, error) {
	var rides []Ride
	err := db.C(COLLECTION).Find(bson.M{}).All(&rides)
	return rides, err
}

// Find a ride by its id
func (r *RidesDAO) FindById(id string) (Ride, error) {
	var ride Ride
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&ride)
	return ride, err
}

// Insert a ride into database
func (r *RidesDAO) Insert(ride Ride) error {
	err := db.C(COLLECTION).Insert(&ride)
	return err
}

// Delete an existing ride
func (r *RidesDAO) Delete(ride Ride) error {
	err := db.C(COLLECTION).Remove(&ride)
	return err
}

// Update an existing movie
// func (m *RidesDAO) Update(ride Ride) error {
// 	err := db.C(COLLECTION).UpdateId(ride.ID, &ride)
// 	return err
// }
