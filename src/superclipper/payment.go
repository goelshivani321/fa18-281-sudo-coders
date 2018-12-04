package main

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
)

// MongoDB Config
var mongodb_server = "localhost:27017"
var mongodb_database = "superclipper"
var mongodb_collection = "payment"

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		formatter.JSON(writer, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

func getPaymentByCardId(formatter *render.Render) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
		//Retrieve the cardId sent as parameter 
        params := mux.Vars(request)
		var cardId string = params["cardId"]

		//Start MongoDB session
        session, error := mgo.Dial(mongodb_server)
        if error != nil {
			formatter.JSON(writer, http.StatusServiceUnavailable, "")
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)
		
		//Find the Document present in MongoDB collection with matching cardId
        var cardPayment = CardPayment{}
        error = collection.Find(bson.M{"cardid" : cardId}).One(&cardPayment)
        if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return		
		}
        formatter.JSON(writer, http.StatusOK, cardPayment)
    }
}

func getPaymentByCardIdPaymentId(formatter *render.Render) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
		//Retrieve cardId and paymentId sent as parameter
        params := mux.Vars(request)
		var cardId string = params["cardId"]
		var paymentId string = params["paymentId"]
		contains := false

		//Start MongoDB session
        session, error := mgo.Dial(mongodb_server)
        if error != nil {
			formatter.JSON(writer, http.StatusServiceUnavailable, "")
			return
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        collection := session.DB(mongodb_database).C(mongodb_collection)

		//Find document in MongoDB collection with matching CardId
        var cardPayment = CardPayment{}
        error = collection.Find(bson.M{"cardid" : cardId}).One(&cardPayment)
        if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return
		}

		//Find payment in MongoDB collection with matching PaymentId
		for _, payment := range cardPayment.Payments {
			if payment.PaymentId == paymentId {
				contains = true
				formatter.JSON(writer, http.StatusOK, payment)
			}
		}
		if(!contains) {
			formatter.JSON(writer, http.StatusNotFound, "")
			return 
		}
    }
}

func updatePaymentByCardIdPaymentId(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//Retrieve cardId and paymentId sent as parameter
		params := mux.Vars(request)
		var cardId string = params["cardId"]
		var paymentId string = params["paymentId"]

		//Create an payment request object
		var paymentRequest = Payment{}
		_ = json.NewDecoder(request.Body).Decode(&paymentRequest)

		//Start MongoDB session
		session, error := mgo.Dial(mongodb_server)
		if error != nil {
			formatter.JSON(writer, http.StatusServiceUnavailable, "")
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		//Find document in MongoDB collection with matching CardId
        var cardPayment = CardPayment{}
		error = collection.Find(bson.M{"cardid" : cardId}).One(&cardPayment)
        if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return
		}

		//Loop through Payments and update payment with values of payment request object
		payments := []Payment{}
		for _, payment := range cardPayment.Payments {
			if payment.PaymentId == paymentId {
				payment = Payment(paymentRequest)
				payment.PaymentId = paymentId
			}
			payments = append(payments,payment)
		}
		cardPayment.SetPayments(payments)

		//Update CardPayment in MongoDB collection
		error = collection.Update(bson.M{"cardid": cardId}, &cardPayment)
		if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return
		}
		formatter.JSON(writer, http.StatusOK, cardPayment)

		//TODO: Hit Cards API '/update/{cardid}/{bal}' of type PUT to update Available Balance in Cards database
	}
}

func createPaymentByCardId(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//Retrieve cardId sent as parameter
		params := mux.Vars(request)
		var cardId string = params["cardId"]

		//Create an payment request object
		var paymentRequest = Payment{}
		_ = json.NewDecoder(request.Body).Decode(&paymentRequest)

		//Start MongoDB session
		session, error := mgo.Dial(mongodb_server)
		if error != nil {
			panic(error)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		//Find document in MongoDB collection with matching CardId
        var cardPayment = CardPayment{}
		error = collection.Find(bson.M{"cardid" : cardId}).One(&cardPayment)
		if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return
		}

		//Add the request Payment object to cardPayment
		paymentRequest.PaymentId = PaymentIdGenerator(cardPayment.Payments)
		cardPayment.Payments = append(cardPayment.Payments, paymentRequest)

		//Update the whole cardPayment document on MongoDB
		error = collection.Update(bson.M{"cardid" : cardId}, cardPayment)
		if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return
		}
		formatter.JSON(writer, http.StatusOK, cardPayment)
	}
}