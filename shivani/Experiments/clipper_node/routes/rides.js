var express = require('express');
var router = express.Router();
var request = require("request");

/* GET rides page. */
router.get('/', function(req, res, next) {

  var options = { method: 'GET',
    url: 'http://localhost:3000/rides/5c035d754bdba90001db1a1f' };

  request(options, function (error, response, body) {
    if (error) throw new Error(error);

    console.log(body);
    data = JSON.parse(body)
    
    res.render('rides',
    {
      rideId: data.id,
      customerID: data.customerID,
      startTime: data.startTime,
      startLocation: data.startLocation,
      rideType: data.rideType
    });
  });

});

/* POST rides page. */
// router.get('/', function(req, res, next) {
//
//   var request = require("request");
//
// var options = { method: 'POST',
//     url: 'http://localhost:3000/rides',
//     headers:
//     body: '{\n\t"customerID" : "1001",\n\t"startTime" : "2018-11-20 10:35",\n\t"startLocation" : "Tasman Drive",\n\t"rideType" : "Light Rail"\n}' };
//
//   request(options, function (error, response, body) {
//     if (error) throw new Error(error);
//
//     console.log(body);
//   });
//
//   res.render('index', { title: 'Express' });
// });

/* GET rides by id */
// router.get('/', function(req, res, next) {
//   var request = require("request");
//
// var options = { method: 'GET',
//     url: 'http://localhost:3000/rides/5c02485f4bdba90001db1a1e',
//
//   request(options, function (error, response, body) {
//     if (error) throw new Error(error);
//
//     console.log(body);
//   });
//
//   res.render('index', { title: body });
// });


module.exports = router;
