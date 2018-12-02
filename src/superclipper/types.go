package main

type Payments struct {
	CardId  string
	Payment string
	// Payments	[]Payment
}

type PaymentTran struct {
	Id 		string 	
	CardId 	string
	Payment string
}

// type Address struct {
// 	FirstName  string
// 	LastName   string
// 	Address1   string
// 	Address2   string
// 	Address3   string
// 	City       string
// 	State      string
// 	PostalCode string
// 	County     string
// 	Country    string
// 	Phone      string
// 	Email      string
// }

// type Payment struct {
// 	CurrencyCode         string
// 	AccountNumber        string
// 	AccountDisplayNumber string
// 	NameOnCard           string
// 	CardExpiryMonth      int
// 	CardExpiryYear       int
// 	Amount               float
// 	ChargeSequence       int
// 	BillingAddress       Address
// 	PaymentType          string
// 	CardTypeId           string
// }
