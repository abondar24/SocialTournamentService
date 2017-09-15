package main

import (
	"github.com/abondar24/SocialTournamentService/api"
	"github.com/abondar24/SocialTournamentService/data"
	"log"
)

func main() {
	ds,err := data.ConnectToBase()
	if err!=nil{
		log.Fatal(err)
	}

	server := api.NewServer(ds)
	server.RunRestServer()

}
