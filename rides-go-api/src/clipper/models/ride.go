package models

import "gopkg.in/mgo.v2/bson"

// Represents a Ride, we use bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Ride struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	CustomerID    string        `bson:"customerID" json:"customerID"`
	StartTime     string        `bson:"startTime" json:"startTime"`
	StartLocation string        `bson:"startLocation" json:"startLocation"`
	RideType      string        `bson:"rideType" json:"rideType"`
}
