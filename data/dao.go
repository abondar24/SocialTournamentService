package data

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

type MySql struct {
	dbInst *sql.DB
}

func ConnectToBase() (*MySql) {
	instance, err := sql.Open("mysql",
		"root:alex21@tcp(172.17.0.2:3306)/social_tournament?charset=utf8")

	if err != nil {
		log.Fatalln(err)
	}

	return &MySql{dbInst: instance,}
}
