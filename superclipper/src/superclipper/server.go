/*
	Gumball API in Go (Version 2)
	Uses MongoDB and RabbitMQ
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	// "github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)


// MongoDB Config
var mongodb_server = "mongodb://admin:cmpe281@10.0.1.130,10.0.1.210,10.0.1.124,10.0.1.241,10.0.1.77"
var mongodb_database = "superclipper"
// var mongodb_payment_collection = "payment"
// var mongodb_cards_collection = "cards"
var mongodb_users_collection = "users"

// RabbitMQ Config
// var rabbitmq_server = "rabbitmq"
// var rabbitmq_port = "5672"
// var rabbitmq_queue = "gumball"
// var rabbitmq_user = "guest"
// var rabbitmq_pass = "guest"

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

    mx.HandleFunc("/users/{userId}", getUsersByUserId(formatter)).Methods("GET")

	mx.HandleFunc("/users/{userId}", updateUsersByUserId(formatter)).Methods("PUT")

	mx.HandleFunc("/users", createUsers(formatter)).Methods("POST")


	// mx.HandleFunc("/payment/{paymentId}", deletePaymentHandler(formatter)).Methods("DELETE")

	// mx.HandleFunc("/order/{id}", gumballOrderStatusHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/order", gumballOrderStatusHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/orders", gumballProcessOrdersHandler(formatter)).Methods("POST")
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

func getUsersByUserId(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        params := mux.Vars(req)
        var userId string = params["userId"]
        session, err := mgo.Dial(mongodb_server)
        if err != nil {
        	fmt.Println("Not connecting to DB")
        	formatter.JSON(w, http.StatusOK, "")
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_users_collection)
        var user Users
        err = c.Find(bson.M{"userid" : userId}).One(&user)
        if err != nil {
        	fmt.Println("User Id not present in mongo")
            formatter.JSON(w, http.StatusOK, "")
        }
        fmt.Println(user)
        formatter.JSON(w, http.StatusOK, user)
    }
}


func updateUsersByUserId(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
        params := mux.Vars(req)
        var userId string = params["userId"]
    	var userReq Users
    	_ = json.NewDecoder(req.Body).Decode(&userReq)
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_users_collection)
        var user Users
        err = c.Find(bson.M{"UserId" : userId}).One(&user)
            if err != nil {
                    log.Fatal(err)
            }
		formatter.JSON(w, http.StatusOK, user)
	}
}


func createUsers(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var userRequest Users
		_ = json.NewDecoder(req.Body).Decode(&userRequest)
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_users_collection)
        err = c.Insert(&userRequest)
        if err != nil {
                panic(err)
        }
	}
}



// func deletePaymentHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {

//         //if to check if empty
//     	params := mux.Vars(req)
//         var paymentId string = params["paymentId"]

//     	fmt.Println("Delete Payment: ", paymentId)

// 		session, err := mgo.Dial(mongodb_server)
//         if err != nil {
//                 panic(err)
//         }
//         defer session.Close()
//         session.SetMode(mgo.Monotonic, true)

//         c := session.DB(mongodb_database).C("paymentTransaction")

//         err = c.Remove(bson.M{"id": paymentId})

//         if err != nil {
//                 log.Fatal(err)
//         }

//         fmt.Println("Payment succesfully deleted", paymentId )
// 		formatter.JSON(w, http.StatusOK)
// 	}
// }


// // API Get Order Status
// func gumballOrderStatusHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		params := mux.Vars(req)
// 		var uuid string = params["id"]
// 		fmt.Println( "Order ID: ", uuid )
// 		if uuid == ""  {
// 			fmt.Println( "Orders:", orders )
// 			var orders_array [] order
// 			for key, value := range orders {
//     			fmt.Println("Key:", key, "Value:", value)
//     			orders_array = append(orders_array, value)
// 			}
// 			formatter.JSON(w, http.StatusOK, orders_array)
// 		} else {
// 			var ord = orders[uuid]
// 			fmt.Println( "Order: ", ord )
// 			formatter.JSON(w, http.StatusOK, ord)
// 		}
// 	}
// }

// // API Process Orders
// func gumballProcessOrdersHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {

// 		// Open MongoDB Session
// 		session, err := mgo.Dial(mongodb_server)
//         if err != nil {
//                 panic(err)
//         }
//         defer session.Close()
//         session.SetMode(mgo.Monotonic, true)
//         c := session.DB(mongodb_database).C(mongodb_collection)

//        	// Get Gumball Inventory
//         var result bson.M
//         err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
//         if err != nil {
//                 log.Fatal(err)
//         }

//  		var count int = result["CountGumballs"].(int)
//         fmt.Println("Current Inventory:", count )

// 		// Process Order IDs from Queue
// 		var order_ids []string = queue_receive()
// 		for i := 0; i < len(order_ids); i++ {
// 			var order_id = order_ids[i]
// 			fmt.Println("Order ID:", order_id)
// 			var ord = orders[order_id]
// 			ord.OrderStatus = "Order Processed"
// 			orders[order_id] = ord
// 			count -= 1
// 		}
// 		fmt.Println( "Orders: ", orders , "New Inventory: ", count)

// 		// Update Gumball Inventory
// 		query := bson.M{"SerialNumber" : "1234998871109"}
//         change := bson.M{"$set": bson.M{ "CountGumballs" : count}}
//         err = c.Update(query, change)
//         if err != nil {
//                 log.Fatal(err)
//         }

// 		// Return Order Status
// 		formatter.JSON(w, http.StatusOK, orders)
// 	}
// }

// // Send Order to Queue for Processing
// func queue_send(message string) {
// 	conn, err := amqp.Dial("amqp://"+rabbitmq_user+":"+rabbitmq_pass+"@"+rabbitmq_server+":"+rabbitmq_port+"/")
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	failOnError(err, "Failed to open a channel")
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		rabbitmq_queue, // name
// 		false,   // durable
