package models

import "gopkg.in/mgo.v2/bson"

type CardHolder struct {
    ID        bson.ObjectId  `bson:"_id",json:"id"`
    FirstName string         `bson:"firstname",json:"firstname"`
    LastName  string         `bson:"lastname",json:"lastname"`
    Address  *Address        `bson:"address",json:"address"`
}

type Address struct {
  City string   `bson:"city",json:"city"`
  State string  `bson:"state",json:"state"`
}
