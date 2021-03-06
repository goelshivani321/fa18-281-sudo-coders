package dao

import (

	"log"

	. "clipper/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RidesDAO struct {
	MongoURI string
}

var db *mgo.Database

const (
	COLLECTION = "rides"
)

// Establish a connection to database
func (r *RidesDAO) Connect() {
	//var mongoURI = "mongodb://username:password@prefix1.mongodb.net,prefix2.mongodb.net,prefix3.mongodb.net/dbName?replicaSet=replName&authSource=admin"

	log.Println("In DAO URI:" + r.MongoURI)
	dialInfo, err := mgo.ParseURL(r.MongoURI)

	log.Println(dialInfo.Database)
	//Below part is similar to above.
	//tlsConfig := &tls.Config{}
	//dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
	//	conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
	//	return conn, err
	//}
	session, err := mgo.DialWithInfo(dialInfo)
	log.Println("In DAO Session:", session)
	//session, err := mgo.Dial(r.Server)

	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(dialInfo.Database)
	log.Println("In DAO Database:", dialInfo.Database)
	//db = session.DB(r.Database)
}

// Find list of rides
func (r *RidesDAO) FindAll() ([]Ride, error) {
	var rides []Ride
	err := db.C(COLLECTION).Find(bson.M{}).All(&rides)
	return rides, err
}

// Find list of rides by CustomerID
func (r *RidesDAO)  FindAllByCustomerId(cid string) ([]Ride, error){
	var rides []Ride
	err := db.C(COLLECTION).Find(bson.M{"customerID": cid}).All(&rides)
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
