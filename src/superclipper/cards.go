
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

		fmt.Println("inside update")
		params := mux.Vars(req)
		//var uuid string = params["name"]
		var cardid string = params["cardid"]
		var bal string = params["bal"]

		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        query := bson.M{"id" : cardid}
        change := bson.M{"$set": bson.M{ "bal" : bal}}
        err = c.Update(query, change)
        if err != nil {
                log.Fatal(err)
        }
       	var result bson.M
        err = c.Find(bson.M{"id" : cardid}).One(&result)
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
		//ids,err := strconv.Atoi(tid)


        

		var mybalance string = params["mybal"]
		var myexp string = params["myexp"]
		//var email string = params["email"]
		
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

        err = c.Insert(&Cards{Id: tid, Bal: mybalance, Expiry:myexp})

		if err != nil {
			panic(err)
		}

	}
}




func readbyid(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("inside read")
		params := mux.Vars(req)
		//var uuid string = params["name"]
		var sid string = params["cardid"]


		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Find(bson.M{"id" : sid}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Data from server is :", result )
		formatter.JSON(w, http.StatusOK, result)



	}
}





func delbyid(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("inside del")
		params := mux.Vars(req)
		//var uuid string = params["name"]
		var sid string = params["cardid"]


		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Remove(bson.M{"id" : sid})
        if err != nil {
                log.Fatal(err)
        }
        
		formatter.JSON(w, http.StatusOK, result)


	}
}


