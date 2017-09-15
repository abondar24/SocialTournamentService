package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySql struct {
	dbInst *sql.DB
}

func ConnectToBase() (*MySql, error) {
	instance, err := sql.Open("mysql",
		"root:alex21@tcp(172.17.0.2:3306)/social_tournament?charset=utf8")

	if err != nil {
		return nil, err
	}

	return &MySql{dbInst: instance}, nil
}

func (ds *MySql) GetPlayerById(playerId int64) (*Player, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM player where id=%v", playerId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		return &Player{}, err
	}

	player := Player{}

	defer stmt.Close()

	rows, err := stmt.Query()

	for rows.Next() {
		rows.Scan(&player.Id, &player.Name, &player.Balance, &player.BackId)
		//if err != nil {
		//	return nil,err
		//}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &player, err

}

func (ds *MySql) CreateNewPlayer(p *Player) (int64, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return int64(0), err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO player(name,balance) VALUES('%v',%v)", p.Name, p.Balance)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return int64(0), err
	}

	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		return int64(0), err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return int64(0), err
	}

	err = tx.Commit()
	if err != nil {
		return int64(0), err
	}

	return int64(id), err
}

func (ds *MySql) UpdatePlayerBalance(player_id int64, sum int64, charge bool) error {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := ""

	if !charge {
		query = fmt.Sprintf("UPDATE player set balance = balance + %v where id=%v", sum, player_id)
	} else {
		query = fmt.Sprintf("UPDATE player set balance = balance - %v where id=%v", sum, player_id)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

func (ds *MySql) GetBalanceForPlayer(playerId int64) (int, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT player.balance FROM player where id=%v", playerId)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return 0, err
	}

	balance := 0

	defer stmt.Close()
	rows, err := stmt.Query()
	for rows.Next() {
		err := rows.Scan(&balance)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return balance, err

}

func (ds *MySql) ClearDB() error {

	err := ds.TruncateSingleTable("backer")
	err = ds.TruncateSingleTable("tournament_player")
	err = ds.TruncateSingleTable("player")
	err = ds.TruncateSingleTable("tournament")

	return err
}

func (ds *MySql) TruncateSingleTable(table string) error {

	tx, err := ds.dbInst.Begin()

	defer tx.Rollback()

	query := fmt.Sprintf("DELETE  FROM %v", table)

	stmt, err := tx.Prepare(query)

	defer stmt.Close()
	_, err = stmt.Exec()

	err = tx.Commit()

	return err
}

func (ds *MySql) DropDatabase(dbName string) error {

	tx, err := ds.dbInst.Begin()

	defer tx.Rollback()

	query := fmt.Sprintf("DROP DATABASE %v", dbName)

	stmt, err := tx.Prepare(query)

	defer stmt.Close()
	_, err = stmt.Exec()

	return err
}
