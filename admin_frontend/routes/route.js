let express = require('express');
let router = express.Router();
let request = require('request');

let baseURI = "http://localhost:8080/v2";


router.get("/status",function (req, res, next) {
    request.get(baseURI+"/",function (err, resp, body) {

        if (!err){
            res.send({msg:JSON.parse(body),code:resp.statusCode});
        } else {
            console.log(err)
        }

    })
});


router.post("/announce_tournament",function (req, res, next) {
    let options = {
        qs: {
            name: req.query.name,
            deposit: parseInt(req.query.deposit)
        }
    };


    request.post(baseURI+"/announce_tournament",options,function (err, resp, body) {

        if (!err){
            res.send({msg:JSON.parse(body),code:resp.statusCode});
        } else {
            console.log(err)
        }

    })
});

router.get("/balance/:id",function (req, res, next) {
    request.get(baseURI+"/balance/"+req.params.id,function (err, resp, body) {

        if (!err){
            res.send({msg:JSON.parse(body),code:resp.statusCode});
        } else {
            console.log(err)
        }

    })
});


router.get("/get_players",function (req, res, next) {
    request.get(baseURI+"/get_players",function (err, resp, body) {

        if (!err){
            res.send({msg:JSON.parse(body),code:resp.statusCode});
        } else {
            console.log(err)
        }

    })
});


router.get("/get_tournaments",function (req, res, next) {
    request.get(baseURI+"/get_tournaments",function (err, resp, body) {

        if (!err){
            res.send({msg:JSON.parse(body),code:resp.statusCode});
        } else {
            console.log(err)
        }

    })
});

router.get("/get_players_tournament/:id",function (req, res, next) {
    request.get(baseURI+"/get_players_tournament/"+req.params.id,function (err, resp, body) {

        if (!err){
            res.send({msg:JSON.parse(body),code:resp.statusCode});
        } else {
            console.log(err)
        }

    })
});


router.get("/result_tournament/:id",function (req, res, next) {
    request.get(baseURI+"/result_tournament/"+req.params.id,function (err, resp, body) {

        if (!err){
            res.send({msg:JSON.parse(body),code:resp.statusCode});
        } else {
            console.log(err)
        }

    })
});

router.put("/update_prizes",function (req, res, next) {
    let options = {
        qs: {
            tournament_id: parseInt(req.query.tournament_id),
            player_id: parseInt(req.query.player_id),
            prize: parseInt(req.query.prize)
        }
    };


    request.put(baseURI+"/update_prizes",options,function (err, resp, body) {

        if (!err){
            res.send({code:resp.statusCode});
        } else {
            console.log(err)
        }

    })
});



module.exports = router;