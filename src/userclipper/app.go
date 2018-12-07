package main

import (
      "encoding/json"
       "log"
       "net/http"

       "gopkg.in/mgo.v2/bson"

       "github.com/gorilla/mux"
       . "github.com/vkoul07/restapi/config"
       . "github.com/vkoul07/restapi/dao"
       . "github.com/vkoul07/restapi/models"
)

var config = Config{}
var dao = CardHoldersDAO{
	
// GET cardholder by ID
func FindCardHoldersEndPoint(w http.ResponseWriter, r *http.Request) {
      params := mux.Vars(r)
      cardholder, err := dao.FindById(params["emailid"])
       if err != nil {
          respondWithError(w, http.StatusBadRequest,"CardHolder not present")
           return
        }
        respondWithJson(w, http.StatusOK, cardholder)
}

//POST new cardholder
func CreateCardHoldersEndPoint(w http.ResponseWriter, r *http.Request){
              defer r.Body.Close()
              var cardholder CardHolder
              if err := json.NewDecoder(r.Body).Decode(&cardholder);
              err != nil {
                respondWithError(w, http.StatusBadRequest, "CardHolder cannot be created")
                return
              }
              cardholder.ID = bson.NewObjectId()
              if err := dao.Insert(cardholder);
              err != nil {
                respondWithError(w, http.StatusInternalServerError, err.Error())
                return
              }
              respondWithJson(w, http.StatusCreated,cardholder)
}
//PUT update a current cardholder
func UpdateCardHoldersEndPoint(w http.ResponseWriter, r *http.Request){
        defer r.Body.Close()
        var cardholder CardHolder
        if err := json.NewDecoder(r.Body).Decode(&cardholder);
        err != nil {
          respondWithError(w , http.StatusInternalServerError, err.Error())
          return
        }
        respondWithJson(w, http.StatusOK, map[string]string{"result":"Done"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
         respondWithJson(w, code, map[string]string{"error":msg})
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

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
        r := mux.NewRouter()
        r.HandleFunc("/cardholders/{emailid}",FindCardHoldersEndPoint ).Methods("GET")
        r.HandleFunc("/cardholders",CreateCardHoldersEndPoint).Methods("POST")
	r.HandleFunc("/cardholders",UpdateCardHoldersEndPoint).Methods("PUT")
        if err := http.ListenAndServe(":3000", r);
        err != nil{
          log.Fatal(err)
        }
}
