// create superclipper collection
// mongo payments
db.payments.insert({
    "cardid" : "1",
    "payments" : [ 
        {
            "paymentid" : "1",
            "currencycode" : "USD",
            "cardnumber" : "8765123409341987",
            "nameoncard" : "Sanjay Nag",
            "cardexpirymonth" : "9",
            "cardexpiryyear" : "2020",
            "amount" : "34.6",
            "paymenttype" : "Credit Card",
            "cardtypeid" : "Visa"
        }, 
        {
            "paymentid" : "2",
            "currencycode" : "INR",
            "cardnumber" : "987654321098123",
            "nameoncard" : "Sanjay Nag",
            "cardexpirymonth" : "11",
            "cardexpiryyear" : "2021",
            "amount" : "74.6",
            "paymenttype" : "Debit Card",
            "cardtypeid" : "Master Card"
        }, 
        {
            "paymentid" : "3",
            "currencycode" : "INR",
            "cardnumber" : "7453453875878983",
            "nameoncard" : "Sanjay Nag ",
            "cardexpirymonth" : "11",
            "cardexpiryyear" : "2022",
            "amount" : "4.6",
            "paymenttype" : "Debit Card",
            "cardtypeid" : "Visa"
        }
    ]
});

db.payments.insert(
{
    "cardid" : "2",
    "payments" : [ 
        {
            "paymentid" : "1",
            "currencycode" : "USD",
            "cardnumber" : "4354129087123409",
            "nameoncard" : "Nachiket Wattamwar",
            "cardexpirymonth" : "1",
            "cardexpiryyear" : "2019",
            "amount" : "96.69",
            "paymenttype" : "Stash Unstash",
            "cardtypeid" : "Shim Paul"
        }, 
        {
            "paymentid" : "4",
            "currencycode" : "USD",
            "cardnumber" : "1230980912321672",
            "nameoncard" : "Sanjay Nag",
            "cardexpirymonth" : "9",
            "cardexpiryyear" : "5050",
            "amount" : "23.0",
            "paymenttype" : "Card",
            "cardtypeid" : "Visa"
        }
    ]
});

db.payments.insert(
{
    "cardid" : "3",
    "payments" : [ 
        {
            "paymentid" : "1",
            "currencycode" : "USD",
            "cardnumber" : "245345345345457",
            "nameoncard" : "Sanjay Nag",
            "cardexpirymonth" : "9",
            "cardexpiryyear" : "2022",
            "amount" : "34.6",
            "paymenttype" : "Credit Card",
            "cardtypeid" : "1"
        }, 
        {
            "paymentid" : "2",
            "currencycode" : "USD",
            "cardnumber" : "9876129045631830",
            "nameoncard" : "Sanjay Nag",
            "cardexpirymonth" : "1",
            "cardexpiryyear" : "2023",
            "amount" : "346",
            "paymenttype" : "Credit Card",
            "cardtypeid" : "1"
        }, 
        {
            "paymentid" : "3",
            "currencycode" : "USD",
            "cardnumber" : "3451097834126546",
            "nameoncard" : "Paul Nguyen",
            "cardexpirymonth" : "5",
            "cardexpiryyear" : "5050",
            "amount" : "230",
            "paymenttype" : "Debit Card",
            "cardtypeid" : "Visa"
        }, 
        {
            "paymentid" : "4",
            "currencycode" : "USD",
            "cardnumber" : "4234345678563453",
            "nameoncard" : "Sanjay Nag",
            "cardexpirymonth" : "9",
            "cardexpiryyear" : "2050",
            "amount" : "234",
            "paymenttype" : "Debit Card",
            "cardtypeid" : "Visa"
        }, 
        {
            "paymentid" : "5",
            "currencycode" : "USD",
            "cardnumber" : "245345345345457",
            "nameoncard" : "Sanjay Nag",
            "cardexpirymonth" : "9",
            "cardexpiryyear" : "2054",
            "amount" : "30",
            "paymenttype" : "Credit Card",
            "cardtypeid" : "Visa"
        }, 
        {
            "paymentid" : "6",
            "currencycode" : "USD",
            "cardnumber" : "3451097834126546",
            "nameoncard" : "Sanjay Nag",
            "cardexpirymonth" : "9",
            "cardexpiryyear" : "2020",
            "amount" : "20",
            "paymenttype" : "Credit Card",
            "cardtypeid" : "Visa"
        }
    ]
});

db.payments.insert(
{
    "cardid" : "4",
    "payments" : [ 
        {
            "paymentid" : "1",
            "currencycode" : "INR",
            "cardnumber" : "245345345345457",
            "nameoncard" : "Sanjay Nag",
            "cardexpirymonth" : "1",
            "cardexpiryyear" : "2020",
            "amount" : "634",
            "paymenttype" : "Credit Card",
            "cardtypeid" : "1"
        }, 
        {
            "paymentid" : "2",
            "currencycode" : "USD",
            "cardnumber" : "2342456587651234",
            "nameoncard" : "Nachiket Wattamwar",
            "cardexpirymonth" : "6",
            "cardexpiryyear" : "2020",
            "amount" : "87.2",
            "paymenttype" : "Credit Card",
            "cardtypeid" : "Visa"
        }
    ]
});