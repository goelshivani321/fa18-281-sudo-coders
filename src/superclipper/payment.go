package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "mongodb"
var mongodb_database = "superclipper"
var mongodb_collection = "payment"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payment", gumballHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payment", gumballUpdateHandler(formatter)).Methods("PUT")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// API Gumball Machine Handler
func gumballHandler(formatter *render.Render) http.HandlerFunc {
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

// API Update Gumball Inventory
func gumballUpdateHandler(formatter *render.Render) http.HandlerFunc {
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
