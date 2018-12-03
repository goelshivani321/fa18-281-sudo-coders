package dao

import (
        "log"

        . "github.com/vkoul07/restapi/models"
        mgo "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type CardHoldersDAO struct {
         Server string
         Database string
}

var db *mgo.Database

const (
       COLLECTION = "cardholders"
)

//Establish connection to the database
func (m *CardHoldersDAO) Connect(){
        session, err := mgo.Dial(m.Server)
        if err != nil {
               log.Fatal(err)
        }
        db = session.DB(m.Database)
}

// Find list of cardholders
func (m *CardHoldersDAO) FindAll() ([]CardHolder,error) {
      var cardholders []CardHolder
      err := db.C(COLLECTION).Find(bson.M{}).All(&cardholders)
      return cardholders,err
}

// Find a cardholder by its id
func(m *CardHoldersDAO) FindById(id string)(CardHolder,error){
     var cardholder CardHolder
     err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&cardholder)
     return cardholder,err
}

//Insert a cardholder
func (m *CardHoldersDAO) Insert(cardholder CardHolder) error {
       err := db.C(COLLECTION).Insert(&cardholder)
       return err
  }

  //Delete an existing cardholder
  func (m *CardHoldersDAO) Delete(cardholder CardHolder) error {
       err := db.C(COLLECTION).Remove(&cardholder)
       return err
  }

  //Update an existing cardholder
  func (m *CardHoldersDAO) Update(cardholder CardHolder) error {
      err := db.C(COLLECTION).UpdateId(cardholder.ID, &cardholder)
      return err
  }
