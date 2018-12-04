package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	newServer := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	newServer.UseHandler(mx)
	return newServer
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payments/cardId/{cardId}", getPaymentByCardId(formatter)).Methods("GET")
	mx.HandleFunc("/payments/cardId/{cardId}/paymentId/{paymentId}", getPaymentByCardIdPaymentId(formatter)).Methods("GET")
	// mx.HandleFunc("/payments/cardId/{cardId}/paymentId/{paymentId}", updatePaymentByCardIdPaymentId(formatter)).Methods("PUT")
	// mx.HandleFunc("/payments/cardId/{cardId}/payment", createPaymentByCardId(formatter)).Methods("POST")
	// mx.HandleFunc("/payments/cardId/{cardId}/payment/{paymentId}", deletePaymentByCardIdPaymentId(formatter)).Methods("DELETE")
}

// // Helper Functions
// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 		panic(fmt.Sprintf("%s: %s", msg, err))
// 	}
// }
