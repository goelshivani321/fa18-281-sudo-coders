package main

import (
	"strconv"
)

type CardPayment struct {
	CardId  	string `json: "id" bson: "_id,omitempty"`
	Payments 	[]Payment `json: "payments" bson: "payments"`
}

type Payment struct {
	PaymentId 				string `json: "paymentid" bson: "paymentid"`
	CurrencyCode			string `json: "currencycode" bson: "currencycode"`
	CardNumber				string `json: "cardnumber" bson: "cardnumber"`
	NameOnCard				string `json: "nameoncard" bson: "nameoncard"`
	CardExpiryMonth			string `json: "cardexpirymonth" bson: "cardexpirymonth"`
	CardExpiryYear			string `json: "cardexpiryyear" bson: "cardexpiryyear"`
	Amount	 				string `json: "amount" bson: "amount"`
	PaymentType				string `json: "paymenttype" bson: "paymenttype"`
	CardType				string `json: "cardtype" bson: "cardtype"`
}

type User struct {
	UserName  	string `json: "username" bson: "username"`
	Password 	string `json: "password" bson: "password"`
}


// SetPayment receives a pointer to CardPayment so it can modify it
func (cardPayment *CardPayment) SetPayments(payments []Payment) {
    cardPayment.Payments = payments
}

// PaymentIdGenerator generates the next PaymentMethodId
func PaymentIdGenerator(payments []Payment) string {
	return strconv.Itoa(len(payments) + 1)
}