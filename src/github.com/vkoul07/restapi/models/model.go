package models

import "gopkg.in/mgo.v2/bson"

type CardHolder struct {
    Id        bson.ObjectId  `bson:"_id",json:"id"`
    FirstName string         `bson:"firstname",json:"firstname"`
    LastName  string         `bson:"lastname",json:"lastname"`
    Address   string         `bson:"address",json:"address"`
    cardinfo                 `bson:"address",json:"address"`
    balance                  `bson:"address",json:"address"`
}
