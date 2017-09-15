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
		return nil, err
	}

	player := &Player{}

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

	return player, err

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

func (ds *MySql) CreateNewTournament(t *Tournament) (int64, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return int64(0), err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO tournament(name,deposit) VALUES('%v',%v)", t.Name, t.Deposit)
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

func (ds *MySql) GetTournamentById(tournamentId int64) (*Tournament, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM tournament where id=%v", tournamentId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		return nil, err
	}

	tournament := &Tournament{}

	defer stmt.Close()

	rows, err := stmt.Query()

	for rows.Next() {
		rows.Scan(&tournament.Id, &tournament.Name, &tournament.Deposit)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return tournament, err

}

func (ds *MySql) AddPlayerToTournament(tp *TournamentPlayer) (int64, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return int64(0), err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO tournament_player(player_id,tournament_id) VALUES('%v',%v)", tp.PlayerId, tp.TournamentId)
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

func (ds *MySql) GetTournamentPlayersByTournamentId(tournamentId int64) (*[]TournamentPlayer, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM tournament_player where tournament_id=%v", tournamentId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	players := make([]TournamentPlayer, len(cols))

	for rows.Next() {
		tp := TournamentPlayer{}
		rows.Scan(&tp.Id, &tp.TournamentId, &tp.PlayerId, &tp.Prize)
		players = append(players, tp)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &players, err
}

func (ds *MySql) UpdateResultOfTournamentForPlayer(tp *TournamentPlayer) error {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("UPDATE tournament_player SET price=%v WHERE tournament_id=%v AND player_id=%v ",
		tp.Prize, tp.TournamentId, tp.PlayerId)

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

func (ds *MySql) GetTournamentWinners(tournamentId int64) (*[]TournamentPlayer, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM tournament_player where tournament_id=%v AND prize>0", tournamentId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	players := make([]TournamentPlayer, len(cols))

	for rows.Next() {
		tp := TournamentPlayer{}
		rows.Scan(&tp.Id, &tp.TournamentId, &tp.PlayerId, &tp.Prize)
		players = append(players, tp)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &players, err
}

func (ds *MySql) BackPlayerForTournament(b *Backer) (int64, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return int64(0), err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO backer(player_id,backer_id,sum) VALUES(%v,%v,%v)", b.PlayerId, b.BackerId, b.Sum)
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

func (ds *MySql) GetPlayerBackers(playerId int64) (*[]Backer, error) {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM backer where player_id=%v ", playerId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	backers := make([]Backer, len(cols))

	for rows.Next() {
		b := Backer{}
		rows.Scan(&b.Id, &b.PlayerId, &b.BackerId, &b.Sum)
		backers = append(backers, b)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &backers, err
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
