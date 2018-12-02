package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "mongodb"
var mongodb_database = "superclipper"
var mongodb_collection = "payment"


// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}


func paymentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var result bson.M
		err = c.Find(bson.M{"CardId": "5018787760363113049"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Gumball Machine:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}


func updatePaymentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var payments Payments
		_ = json.NewDecoder(req.Body).Decode(&payments)
		fmt.Println("Update Payment To: ", payments.Payment)
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		query := bson.M{"CardId": "5018787760363113049"}
		change := bson.M{"$set": bson.M{"Payment": payments.Payment}}
		err = c.Update(query, change)
		if err != nil {
			log.Fatal(err)
		}
		var result bson.M
		err = c.Find(bson.M{"CardId": "5018787760363113049"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Payments:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}


func newPaymentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var paymentTranReq PaymentTran
		_ = json.NewDecoder(req.Body).Decode(&paymentTranReq)	

		fmt.Println("Create PaymentTran To: ", paymentTranReq)

		uuid, _ := uuid.NewV4()
    	var paymentTranPer = PaymentTran {
					Id: uuid.String(),            		
					CardId: paymentTranReq.CardId,
					Payment: paymentTranReq.Payment,
		}

		fmt.Println( "paymentTran1: ", paymentTranPer )
		formatter.JSON(w, http.StatusOK, paymentTranPer)


		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C("paymentTransaction")

        err = c.Insert(&paymentTranPer)

        if err != nil {
                panic(err)
        }

        //TODO: RabbitMQ

	}
}