var express = require('express');
var router = express.Router();
var request = require("request");

/* GET rides page. */
router.get('/', function(req, res, next) {

  var options = { method: 'GET',
    url: 'http://localhost:3000/read/1' };

  request(options, function (error, response, body) {
    if (error) throw new Error(error);

    console.log(body);
    data = JSON.parse(body)
    
    res.render('card details',
    {
      cardId: data.id,
      balance: data.bal,
      expiryDate: data.exp
    });
  });

});