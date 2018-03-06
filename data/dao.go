package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type MySql struct {
	dbInst *sql.DB
}

func ConnectToBase() (*MySql, error) {
	instance, err := sql.Open("mysql",
		"root:alex21@tcp(db:3306)/social_tournament?charset=utf8")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &MySql{dbInst: instance}, nil
}

func ConnectToTestBase() (*MySql, error) {
	instance, err := sql.Open("mysql",
		"root:alex21@tcp(localhost:3306)/social_tournament?charset=utf8")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &MySql{dbInst: instance}, nil
}

func (ds *MySql) BeginTx() (*sql.Tx, error) {
	tx, err := ds.dbInst.Begin()
	if err != nil {
		log.Println(err.Error())

	}

	return tx, err
}

func (ds *MySql) GetAllPlayers(tx *sql.Tx) (*[]Player, error) {
	query := fmt.Sprintf("SELECT * FROM player")
	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	players := make([]Player, 0)

	for rows.Next() {
		b := Player{}
		rows.Scan(&b.Id, &b.Name, &b.Points)
		players = append(players, b)
	}

	return &players, err
}

func (ds *MySql) GetPlayerById(playerId int64, tx *sql.Tx) (*Player, error) {
	query := fmt.Sprintf("SELECT * FROM player where id=%v", playerId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	player := &Player{}

	defer stmt.Close()

	rows, err := stmt.Query()

	for rows.Next() {
		rows.Scan(&player.Id, &player.Name, &player.Points)

	}

	return player, err

}

func (ds *MySql) GetPlayersByIds(playerIds *[]int64, tx *sql.Tx) (*[]Player, error) {
	arr := ""
	for i, id := range *playerIds {

		s := strconv.FormatInt(id, 10)
		arr += s
		if i != len(*playerIds)-1 {
			arr += ","
		}

	}
	query := fmt.Sprintf("SELECT * FROM player where id in (%v )", arr)
	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	backers := make([]Player, 0)

	for rows.Next() {
		b := Player{}
		rows.Scan(&b.Id, &b.Name, &b.Points)
		backers = append(backers, b)
	}

	return &backers, err

}

func (ds *MySql) CreateNewPlayer(p *Player, tx *sql.Tx) (int64, error) {

	query := fmt.Sprintf("INSERT INTO player(name,points) VALUES('%v',%v)", p.Name, p.Points)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return int64(0), err
	}

	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return int64(0), err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return int64(0), err
	}

	return int64(id), err
}

func (ds *MySql) UpdatePlayerBalance(playerId int64, sum int, charge bool, tx *sql.Tx) error {
	query := ""

	if !charge {
		query = fmt.Sprintf("UPDATE player set points = points + %v where id=%v", sum, playerId)
	} else {
		query = fmt.Sprintf("UPDATE player set points = points - %v where id=%v", sum, playerId)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (ds *MySql) UpdatePlayersBalance(playerIds *[]int64, sum int, tx *sql.Tx) error {
	ids := ""
	for i, id := range *playerIds {

		s := strconv.FormatInt(id, 10)
		ids += s
		if i != len(*playerIds)-1 {
			ids += ","
		}

	}

	query := fmt.Sprintf("UPDATE player SET points = points - %v WHERE id in (%v) ", sum, ids)

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (ds *MySql) GetBalanceForPlayer(playerId int64, tx *sql.Tx) (int, error) {

	query := fmt.Sprintf("SELECT player.points FROM player where id=%v", playerId)
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
			log.Println(err.Error())
			return 0, err
		}
	}

	return balance, err

}

func (ds *MySql) CreateNewTournament(t *Tournament, tx *sql.Tx) (int64, error) {
	query := fmt.Sprintf("INSERT INTO tournament(name,deposit) VALUES('%v',%v)", t.Name, t.Deposit)
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return int64(0), err
	}

	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return int64(0), err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return int64(0), err
	}

	return int64(id), err
}

func (ds *MySql) GetAllTournamets(tx *sql.Tx) (*[]Tournament, error) {
	query := fmt.Sprintf("SELECT * FROM tournament")
	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	tournaments := make([]Tournament, 0)

	for rows.Next() {
		t := Tournament{}
		rows.Scan(&t.Id, &t.Name, &t.Deposit)
		tournaments = append(tournaments, t)
	}

	return &tournaments, err
}

func (ds *MySql) GetTournamentById(tournamentId int64, tx *sql.Tx) (*Tournament, error) {
	query := fmt.Sprintf("SELECT * FROM tournament where id=%v", tournamentId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	tournament := &Tournament{}

	defer stmt.Close()

	rows, err := stmt.Query()

	for rows.Next() {
		rows.Scan(&tournament.Id, &tournament.Name, &tournament.Deposit)
	}

	return tournament, err

}

func (ds *MySql) AddPlayerToTournament(tp *TournamentPlayer, tx *sql.Tx) (int64, error) {

	query := fmt.Sprintf("INSERT INTO tournament_player(player_id,tournament_id,prize) VALUES(%v,%v,%v)",
		tp.PlayerId, tp.TournamentId, tp.Prize)
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return int64(0), err
	}

	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		return int64(0), err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return int64(0), err
	}

	return int64(id), err

}

func (ds *MySql) GetPlayerIdsByTournament(tournamentId int64, tx *sql.Tx) (*[]int64, error) {
	query := fmt.Sprintf("SELECT player_id FROM tournament_player where tournament_id=%v", tournamentId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	//it's a change in Go 1.8
	//if you see DB Example from GoBase it's different a little
	players := make([]int64, 0)

	id := int64(0)
	for rows.Next() {
		rows.Scan(&id)
		players = append(players, id)
	}

	return &players, err
}

func (ds *MySql) GetPlayersByTournament(tournamentId int64, tx *sql.Tx) (*[]Player, error) {
	//alternative variant: SELECT * FROM player p WHERE p.id IN
	// (SELECT player_id FROM tournament_player tp WHERE tp.tournament_id=%v) ORDER BY p.id ASC
	query := fmt.Sprintf("SELECT p.id,p.name,p.points FROM player p JOIN tournament_player tp "+
		"ON p.id=tp.player_id WHERE tp.tournament_id=%v", tournamentId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	players := make([]Player, 0)

	for rows.Next() {
		p := Player{}
		rows.Scan(&p.Id, &p.Name, &p.Points)
		players = append(players, p)
	}

	return &players, err
}

func (ds *MySql) GetTournamentPlayerIdFromTournament(player int64, tournament int64, tx *sql.Tx) (int64, error) {

	query := fmt.Sprintf("SELECT player_id FROM tournament_player WHERE tournament_id=%v AND player_id=%v", tournament, player)
	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return int64(0), err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	id := int64(0)
	for rows.Next() {
		rows.Scan(&id)
	}

	return id, err
}

func (ds *MySql) SetPlayerPrize(tp *TournamentPlayer, tx *sql.Tx) error {
	query := fmt.Sprintf("UPDATE tournament_player SET prize=%v WHERE tournament_id=%v AND player_id=%v ",
		tp.Prize, tp.TournamentId, tp.PlayerId)

	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (ds *MySql) GetTournamentWinners(tournamentId int64, tx *sql.Tx) (*[]TournamentPlayer, error) {
	query := fmt.Sprintf("SELECT * FROM tournament_player where tournament_id=%v AND prize>0", tournamentId)
	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	players := make([]TournamentPlayer, 0)

	for rows.Next() {
		tp := TournamentPlayer{}
		rows.Scan(&tp.Id, &tp.TournamentId, &tp.PlayerId, &tp.Prize)
		players = append(players, tp)
	}

	return &players, err
}

func (ds *MySql) BackPlayerForTournament(backers *[]Backer, tx *sql.Tx) error {
	query := fmt.Sprintf("INSERT INTO backer(player_id,backer_id,sum) VALUES ")

	vals := ""
	for i, b := range *backers {
		vals += fmt.Sprintf("(%v,%v,%v)", b.PlayerId, b.BackerId, b.Id)
		if i != len(*backers)-1 {
			vals += ","
		}
	}

	query += vals

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err

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
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("DELETE  FROM %v", table)

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (ds *MySql) DropDatabase(dbName string) error {

	tx, err := ds.dbInst.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer tx.Rollback()

	query := fmt.Sprintf("DROP DATABASE %v", dbName)

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}
