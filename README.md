# Small backend service on Go

# Intro
It's a backend for a FAKE gaming service.
This service has been developed purely for a demo purpose.There is a lot of stuff to make it production ready.

I wanted to show the following stuff:
 - transactions in Go
 - gorilla mux and parameters
 - deploying Go app in Docker
 - docker-compose usage

# Idea

The idea is that people can use service and take part in different games. Every player has it's own balance of points 
on which he can take part in different games. 
Every game has it's entry deposit and if player doesn't have enough points to play 
other participants of the same game can back him up with money. 
In such case they get the money back + % of prize equal to 
% entry deposit he had backed the player .


# Implementation

The service implemented as REST API and stores data in relational database
You can get the idea of database structure if you look at db.sql file

REST API has the following set of methods

- Charge player
```yaml
GET /take?player_id=P1&points=300
Takes points from player's balance
Method returns the following response codes:
    200 – charging completed
    404 – player not found
    402 - not enough points to charge
    500 – internal error

```

- Add points to player
```yaml
Ads points to player's balance
GET /fund?player_id=P2&points=300
Method returns the following response codes:
    200 – points added
    404 – player not found
    500 – internal error
```

- Create a new player
```yaml
Creates a new player
GET /add_player?name=Alex&points=300
Method returns the following response codes:
    201 – player created
    500 – internal error
```

- Create a new tournament
```yaml
Creates a new tournament
GET /announceTournament?name=1&deposit=1000
Method returns the following response codes:
    201 – tournament created
    500 – internal error
```

- Add player to tournament and add his backers 
```yaml
Adds a new player to the tournament
GET /joinTournament?tournament_id=1&player_id=P1&backer_id=P2&backer_id=P3
Method returns the following response codes:
    200 – player added 
    401 – backer doesn't take part in the tournament
    402 – player doesn't have enough points
    409 – player is already taking part in the tournament
    404 – player/backer/tournament not found
    500 – internal error

backer_id is optional
```

- Update prizes
```yaml
Updates player's prize
GET /update_prizes?tournament_id=1&player_id=P1&prize=500
Method returns the following response codes:
    200 – prize updated
    500 – internal error
```

- Result tournament
```yaml
Returns results of the tournament
POST /result_tournament?tournament_id=1
Method returns the following response:
    {"tournamentId": "1", "winners": [{"playerId": "P1", "prize": 500}]}
Method returns the following response codes:
    404 – player/backer/tournament not found
    500 – internal error
```

- Balance
```yaml
Returns player's balance
GET /balance?player_id=1
Method returns the following response:
    {"playerId": "P1", "balance": 456}
Method returns the following response codes:
    404 – player not found
    500 – internal error
```

- Reset
```yaml
Clears data from the database
GET /reset

Method returns the following response codes:
    200 – database is reseted successfully 
    500 – internal error
```

# Install and run

- To build a project and use it on your own machine run make install and run ./main
- To build a docker image run docker build -t <name> . If you want to run with compose,change image name in yml

- To run in docker without docker compose:
     
  ```
  create a database container with name db(to be able to link)
  docker run -it --name=<your_name> -p 8080:8080 --link db abondar/socialtournament 
  ```

- To run in docker via docker compose:
  ```yaml
   docker-compose up -d
  ```
If you are not using docker compose you need to deploy MySQL database with root password from db.sql file and link it to container when you run
If you are using docker compose database with empty tables is created automatically

# TODO and issues
- Write a frontend using Vue.JS + Node.js
- fix problem with make test.Currently you can run te