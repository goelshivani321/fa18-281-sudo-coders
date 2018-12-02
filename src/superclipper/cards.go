
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "mongodb"
var mongodb_database = "cards"
var mongodb_collection = "names"




func ping(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"testing"})
	}
}


func read(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Find(bson.M{"Name" : "nachiket"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Data from server is :", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Update Gumball Inventory
func update(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("inside read")
		params := mux.Vars(req)
		//var uuid string = params["name"]
		var email string = params["email"]

		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        query := bson.M{"Name" : "nachiket"}
        change := bson.M{"$set": bson.M{ "email" : email}}
        err = c.Update(query, change)
        if err != nil {
                log.Fatal(err)
        }
       	var result bson.M
        err = c.Find(bson.M{"Name" : "nachiket"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }        
        fmt.Println("my result", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}



func createcard(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {


		fmt.Println("inside create func")
		params := mux.Vars(req)
		//var uuid string = params["name"]
		var tid string = params["ids"]
		ids,err := strconv.Atoi(tid)

		if err != nil {
        // handle error
        }
        ids=ids+1

		var name string = params["name"]
		var surname string = params["surname"]
		var email string = params["email"]
		
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

        err = c.Insert(&Cards{Name: name, Surname: surname, Email:email})

		if err != nil {
			panic(err)
		}

	}
}




func readbyname(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("inside read")
		params := mux.Vars(req)
		//var uuid string = params["name"]
		var name string = params["name"]


		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Find(bson.M{"Name" : name}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Data from server is :", result )
		formatter.JSON(w, http.StatusOK, result)


	}
}

