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

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/payment", getpayments(formatter)).Methods("GET")
	mx.HandleFunc("/ping", ping(formatter)).Methods("GET")
	mx.HandleFunc("/read", read(formatter)).Methods("GET")
	mx.HandleFunc("/update/{email}", update(formatter)).Methods("PUT")
	//mx.HandleFunc("/create/{name}/{surname}/{email}", gumballNewOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/read/{name}", readbyname(formatter)).Methods("GET")


}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}