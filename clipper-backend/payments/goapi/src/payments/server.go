package main

import (
	"os"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/unrolled/render"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// MongoDB Config
var mongodb_server = os.Getenv("MONGODB")
var mongodb_database = os.Getenv("DATABASE")
var mongodb_collection = os.Getenv("COLLECTION")

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
		corsObj := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
        AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Requested-With"},
    })
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.Use(corsObj)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")

	mx.HandleFunc("/payments/cardId/{cardId}/payment", createPaymentByCardId(formatter)).Methods("POST")
	mx.HandleFunc("/payments/cardId/{cardId}/paymentId/{paymentId}", getPaymentByCardIdPaymentId(formatter)).Methods("GET")
	mx.HandleFunc("/payments/cardId/{cardId}/paymentId/{paymentId}", updatePaymentByCardIdPaymentId(formatter)).Methods("PUT")
	mx.HandleFunc("/payments/cardId/{cardId}/payment/{paymentId}", deletePaymentByCardIdPaymentId(formatter)).Methods("DELETE")

	mx.HandleFunc("/payments/save", createCardPayment(formatter)).Methods("POST")
	mx.HandleFunc("/payments/{cardId}", getPaymentByCardId(formatter)).Methods("GET")
	mx.HandleFunc("/payments/{cardId}", deleteCardPayment(formatter)).Methods("DELETE")

}

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


func createCardPayment(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		//Create an payment request object
		var cardPaymentRequest = CardPayment{}
		_ = json.NewDecoder(request.Body).Decode(&cardPaymentRequest)

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
		error = collection.Find(bson.M{"cardid" : cardPaymentRequest.CardId}).One(&cardPayment)

		if error != nil {
			cardPayment.CardId = cardPaymentRequest.CardId
			payments := []Payment{}
			for _, payment := range cardPaymentRequest.Payments {
				payment.PaymentId = PaymentIdGenerator(payments)
				payments = append(payments, payment)
			}
			cardPayment.SetPayments(payments)
		} else {
			payments := cardPayment.Payments
			for _, payment := range cardPaymentRequest.Payments {
				payment.PaymentId = PaymentIdGenerator(payments)
				payments = append(payments, payment)
			}
			cardPayment.SetPayments(payments)
		}

		//Update the whole cardPayment document on MongoDB
		_, error = collection.Upsert(bson.M{"cardid" : cardPayment.CardId}, &cardPayment)
		if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return
		}

		formatter.JSON(writer, http.StatusOK, cardPayment)
	}
}

func deletePaymentByCardIdPaymentId(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//Retrieve cardId and paymentId sent as parameter
		params := mux.Vars(request)
		var cardId string = params["cardId"]
		var paymentId string = params["paymentId"]

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

		//Loop through Payments and find the payment to be deleted
		payments := []Payment{}
		for _, payment := range cardPayment.Payments {
			if payment.PaymentId != paymentId {
				payment.PaymentId = PaymentIdGenerator(payments)
				payments = append(payments,payment)
			}
		}
		cardPayment.SetPayments(payments)

		//Update the whole cardPayment document on MongoDB
		error = collection.Update(bson.M{"cardid" : cardId}, cardPayment)
		if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return
		}
		formatter.JSON(writer, http.StatusOK, cardPayment)
	}
}

func deleteCardPayment(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//Retrieve cardId sent as parameter
		params := mux.Vars(request)
		var cardId string = params["cardId"]

		//Start MongoDB session
		session, error := mgo.Dial(mongodb_server)
		if error != nil {
			panic(error)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		session.SetSafe(&mgo.Safe{})
		collection := session.DB(mongodb_database).C(mongodb_collection)

		//Find document in MongoDB collection with matching CardId
        var cardPayment = CardPayment{}
		error = collection.Find(bson.M{"cardid" : cardId}).One(&cardPayment)
		if error != nil {
			formatter.JSON(writer, http.StatusNotFound, "")
			return
		}

		formatter.JSON(writer, http.StatusOK, "Deleted Successfully")
	}
}