package main

import (
	"github.com/abondar24/SocialTournamentService/api"
	"github.com/abondar24/SocialTournamentService/data"
	"log"
	"github.com/abondar24/SocialTournamentService/blogic"
	"fmt"
)

func main() {
	ds, err := data.ConnectToBase()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	l:= blogic.NewLogic(ds)
	srv := api.NewServer(l)
	srv.RunRestServer()

}
