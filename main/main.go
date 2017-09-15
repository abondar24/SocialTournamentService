package main

import (
	"github.com/abondar24/SocialTournamentService/data"
	"github.com/abondar24/SocialTournamentService/api"

)
func main(){
    ds:=data.ConnectToBase()

    server:= api.NewServer(ds)
    server.RunRestServer()

}
