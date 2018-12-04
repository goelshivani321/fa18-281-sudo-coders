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
	AccountNumber			string `json: "accountnumber" bson: "accountnumber"`
	AccountDisplayNumber	string `json: "accountdisplaynumber" bson: "accountdisplaynumber"`
	NameOnCard				string `json: "nameoncard" bson: "nameoncard"`
	CardExpiryMonth			int `json: "cardexpirymonth" bson: "cardexpirymonth"`
	CardExpiryYear			int `json: "cardexpiryyear" bson: "cardexpiryyear"`
	Amount	 				float64 `json: "amount" bson: "amount"`
	PaymentType				string `json: "paymenttype" bson: "paymenttype"`
	CardTypeId				string `json: "cardtypeid" bson: "cardtypeid"`
}


// SetPayment receives a pointer to CardPayment so it can modify it
func (cardPayment *CardPayment) SetPayments(payments []Payment) {
    cardPayment.Payments = payments
}

// PaymentIdGenerator generates the next PaymentMethodId
func PaymentIdGenerator(payments []Payment) string {
	return strconv.Itoa(len(payments) + 1)
}