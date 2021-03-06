package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
)

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


func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", ping(formatter)).Methods("GET")
	mx.HandleFunc("/read", read(formatter)).Methods("GET")
	mx.HandleFunc("/update/{email}", update(formatter)).Methods("PUT")
	mx.HandleFunc("/create/{name}/{surname}/{email}", createcard(formatter)).Methods("POST")
	mx.HandleFunc("/read/{name}", readbyid(formatter)).Methods("GET")
	mx.HandleFunc("/delete/{cardid}", delbyid(formatter)).Methods("DEL")


}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}