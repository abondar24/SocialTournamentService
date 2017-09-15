package data

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
	"fmt"
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

func (ds *MySql) GetPlayerById(playerId int64) (*Player, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return &Player{},err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM player where id=%v", playerId)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return &Player{},err
	}

	player := Player{}

	rows, err := stmt.Query()
	for rows.Next() {
		err := rows.Scan(&player.Id, &player.Name, &player.Balance, &player.BackId)
		if err != nil {
			return &Player{},err
		}
	}

	err = tx.Commit()
	if err != nil {
		return &Player{},err
	}
	stmt.Close()

	return &player, err

}

func (ds *MySql) CreateNewPlayer(p *Player) (int64, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return int64(0),err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO player(name,balance) VALUES(?,?)")
	stmt, err := tx.Prepare(query)
	if err != nil {
		return int64(0),err
	}

	res, err := stmt.Exec(p.Name,p.Balance)
	if err != nil {
		return int64(0),err
	}


	id, err := res.LastInsertId()
	if err != nil {
		return int64(0),err
	}
	stmt.Close()

	return int64(id), err
}

func (ds *MySql) ClearDB() error {

	err:= ds.TruncateSingleTable("backer")
	err = ds.TruncateSingleTable("tournament_player")
	err = ds.TruncateSingleTable("player")
	err = ds.TruncateSingleTable("tournament")

	return err
}

func (ds *MySql) TruncateSingleTable(table string) error {


	tx, err := ds.dbInst.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	query := fmt.Sprintf("DELETE  FROM %v", table)

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}


	stmt.Close()

	return err
}


func (ds *MySql) DropDatabase(dbName string) error {


	tx, err := ds.dbInst.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	query := fmt.Sprintf("DROP DATABASE %v", dbName)

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}


	stmt.Close()

	return err
}