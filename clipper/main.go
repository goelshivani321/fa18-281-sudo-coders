package main

import (
	"encoding/json"
	"log"
	"net/http"


	"gopkg.in/mgo.v2/bson"

	. "clipper/config"
	. "clipper/dao"
	. "clipper/models"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = RidesDAO{}

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
	if err := dao.Insert(ride); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, ride)
}

// PUT update an existing ride
// func UpdateRideEndPoint(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	var ride Ride
// 	if err := json.NewDecoder(r.Body).Decode(&ride); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	if err := dao.Update(ride); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
// }

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
	config.Read()
	//config.Server = "mongodb"
	//config.Database = "ridesDB"

	//config.Server = os.Getenv("MONGO_SERVER")
	//config.Database = os.Getenv("MONGO_DB")
	//config.MongoURI = os.Getenv("MONGO_URI")
	dao.Server = config.Server
	dao.Database = config.Database
	dao.MongoURI = config.MongoURI
	//dao.Server = "mongodb"
	//dao.Database = "ridesDB"
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	r.HandleFunc("/rides", AllRidesEndPoint).Methods("GET")
	r.HandleFunc("/rides", CreateRideEndPoint).Methods("POST")
	// r.HandleFunc("/rides", UpdateRideEndPoint).Methods("PUT")
	r.HandleFunc("/rides", DeleteRideEndPoint).Methods("DELETE")
	r.HandleFunc("/rides/{id}", FindRideEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
