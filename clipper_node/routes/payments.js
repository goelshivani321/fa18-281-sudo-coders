var express = require('express');
var router = express.Router();
var request = require("request");


/* GET rides page. */
router.get('/', function(req, res, next) {

  var options = { method: 'GET',
    url: 'localhost:3000/payments/cardId/2' };

  request(options, function (error, response, body) {
    if (error) throw new Error(error);

    console.log(body);
    data = JSON.parse(body)
    
    res.render('Payment details',
    {
      CardId: data.CardId,
      Payments: data.Payments
    });
  });

});