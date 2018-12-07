package main

type Users struct {
	UserId			string `bson:"userid" json:"userid"`
	Firstname 		string `bson:"firstname" json:"firstname"`
	Secondname 		string `bson:"secondname" json:"secondname"`
	Emailid 		string `bson:"emailid" json:"emailid"`
	AddressLine1 	string `bson:"addressline1" json:"addressline1"`
	AddressLine2 	string `bson:"addressline2" json:"addressline2"`
	City 			string `bson:"city" json:"city"`
	State 			string `bson:"state" json:"state"`
	PinCode 		string `bson:"pincode" json:"pincode"`
}
