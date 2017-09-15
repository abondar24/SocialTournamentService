package main

import (
	"github.com/abondar24/SocialTournamentService/api"
	"github.com/abondar24/SocialTournamentService/data"
)

func main() {
	ds := data.ConnectToBase()

	server := api.NewServer(ds)
	server.RunRestServer()

}
