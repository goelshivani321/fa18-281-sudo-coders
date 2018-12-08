package main

import (
	"encoding/json"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/mgo.v2/bson"

	. "clipper/config"
	. "clipper/dao"
	. "clipper/models"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = RidesDAO{}
var quDao = QueueDAO{}

const (
	RideDurationInSecs = 30
	TimeFormat         = "Mon Jan 2 15:04:05 MST 2006"
)

//ping API

func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, "pong")
}

// GET list of rides
func AllRidesEndPoint(w http.ResponseWriter, r *http.Request) {
	rides, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, rides)
}

// GET list of rides by customer id
func AllRidesByCustomerEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerId := params["cid"]
	log.Println(customerId)
	rides, err := dao.FindAllByCustomerId(customerId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	msgs, err := quDao.QueueReceive(customerId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(msgs)

	for i, ride := range rides {
		if isRideInMQ(ride.ID.String(), msgs) {
			startTime,_ := time.Parse(TimeFormat, ride.StartTime)
			if time.Since(startTime).Seconds() > RideDurationInSecs {
				log.Println("Marking live ride as finished: " + ride.ID.String())
				rides[i].LiveStatus = "finished"
			} else {
				log.Println("Re-Adding Current Ride ID: " + ride.ID.String() + " to MQ: " + ride.CustomerID)
				if err := quDao.QueueSend(ride.ID.String(), ride.CustomerID); err != nil {
					respondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}
			}

		} else {
			log.Println("Marking ride as finished: " + ride.ID.String())
			rides[i].LiveStatus = "finished"
		}
	}
	
	respondWithJson(w, http.StatusOK, rides)
}

func isRideInMQ(rideId string, msgs []string) bool {
	for _, msg := range msgs {
		if msg == rideId {
			return true
		}
	}
	return false
}

// GET a ride by its ID
func FindRideEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ride, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Ride ID")
		return
	}
	respondWithJson(w, http.StatusOK, ride)
}

// POST a new ride
func CreateRideEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var ride Ride
	if err := json.NewDecoder(r.Body).Decode(&ride); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	ride.ID = bson.NewObjectId()
	ride.StartTime = time.Now().Format(TimeFormat)
	ride.LiveStatus = "live"
	log.Println(ride.StartTime)

	log.Println("Adding Ride ID: " + ride.ID.String() + " to MQ: " + ride.CustomerID)
	if err := quDao.QueueSend(ride.ID.String(), ride.CustomerID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := dao.Insert(ride); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, ride)
}

// DELETE an existing ride
func DeleteRideEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var ride Ride
	if err := json.NewDecoder(r.Body).Decode(&ride); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(ride); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	//config.Read()
	//config.Server = "mongodb"
	//config.Database = "ridesDB"

	config.MongoURI = os.Getenv("MONGO_URI")
	config.RabbitmqServer = os.Getenv("RABBITMQ_SERVER")
	config.RabbitmqPort = os.Getenv("RABBITMQ_PORT")
	config.RabbitmqUser = os.Getenv("RABBITMQ_USER")
	config.RabbitmqPassword = os.Getenv("RABBITMQ_PASSWORD")

	log.Println("URI: " + config.MongoURI)
	dao.MongoURI = config.MongoURI
	quDao.User = config.RabbitmqUser
	quDao.Password = config.RabbitmqPassword
	quDao.Server = config.RabbitmqServer
	quDao.Port = config.RabbitmqPort
	//dao.Server = "mongodb"
	//dao.Database = "ridesDB"
	dao.Connect()
	quDao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	r.HandleFunc("/rides", AllRidesEndPoint).Methods("GET")
	r.HandleFunc("/rides/{cid}", AllRidesByCustomerEndPoint).Methods("GET")
	r.HandleFunc("/rides", CreateRideEndPoint).Methods("POST")
	r.HandleFunc("/rides", DeleteRideEndPoint).Methods("DELETE")
	r.HandleFunc("/ridesid/{id}", FindRideEndpoint).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	if err := http.ListenAndServe(":3000", c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}