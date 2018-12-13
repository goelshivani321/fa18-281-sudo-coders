package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
	"github.com/rs/cors"
)

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
	mx.HandleFunc("/ping", ping(formatter)).Methods("GET")
	mx.HandleFunc("/read", read(formatter)).Methods("GET")
	mx.HandleFunc("/update/{cardid}/{bal}", update(formatter)).Methods("PUT")
	mx.HandleFunc("/create/{ids}/{mybal}/{myexp}", create(formatter)).Methods("POST")
	mx.HandleFunc("/read/{cardid}", readbyid(formatter)).Methods("GET")
	mx.HandleFunc("/delete/{cardid}", delbyid(formatter)).Methods("DEL")


}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
