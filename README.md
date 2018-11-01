# Small service on Go

# Intro
It's a backend for a FAKE gaming service.
This service has been developed purely for a demo purpose.There is a lot of stuff to make it production ready.

I wanted to show the following stuff:
 - transactions in Go
 - gorilla mux and parameters
 - deploying Go app in Docker
 - docker-compose usage
 - generating Swagger-based api description

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

You can find REST API specification on SwaggerHub:
https://swaggerhub.com/apis/abondar/SocialTournamentService/1.0.0
Or you can access it locally by http://localhost:8024/docs

# Install and run

1) To uild a project and use it on your own machine run make install and run ./main (in such case change database hostname from 'db')
2) To run tests you should have a MySqlServer up and running on localhost
3) To Generate Swagger API reference

```yaml
  Install go-swagger
  In api package dir
    
  swagger generate spec -o ./swagger.json
  swagger validate swagger.json
```
4) To run in docker via docker compose:
  ```yaml
   docker-compose build
   docker-compose up -d
  ``` 
In swagger-alt dir located a Dockerfile for swagger-ui generation based on go-swagger docker image. Just for your interest.


Swagger API documentation is available at localhost:8024
# Frontend
Frontend documentation [here](admin_frontend/README.md) 
