var express = require('express');
var router = express.Router();
var request = require('request');

var baseURI = "http://localhost:8080/v2/";


router.get('/status',function (req, res, next) {
    request(baseURI,function (err, resp, body) {

        if (!err){
            res.send({msg:body,code:resp.statusCode});
        } else {
            console.log(err)
        }

    })
});


module.exports = router;