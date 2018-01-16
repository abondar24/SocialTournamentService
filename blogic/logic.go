package blogic

import (
	"database/sql"
	"errors"
	"github.com/abondar24/SocialTournamentService/data"
	"log"
)

const (
	ErrPlayerNotFound            = "player not found"
	ErrTournamentNotFound        = "tournament not found"
	ErrInternalError             = "internal error"
	ErrInsufficientBalance       = "not enough points"
	ErrPlayerAlreadyInTournament = "player is already in tournament"
)

type Logic struct {
	ds *data.MySql
}

func NewLogic(dataSource *data.MySql) *Logic {

	return &Logic{
		dataSource,
	}
}

func (l *Logic) AddPlayer(name string, points int) (int64, error) {
	p := &data.Player{
		Name:   name,
		Points: points,
	}

	tx, err := l.ds.BeginTx()
	if err != nil {
		log.Println(err.Error())
	}

	defer tx.Rollback()

	pid, err := l.ds.CreateNewPlayer(p, tx)
	if err != nil {
		log.Println(err.Error())
		err = errors.New(ErrInternalError)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}

	return pid, err
}

func (l *Logic) Take(playerId int64, points int) error {
	tx, err := l.ds.BeginTx()
	if err != nil {
		log.Println(err.Error())
	}

	defer tx.Rollback()

	_, err = l.checkPlayer(playerId, tx)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = l.checkBalanceForCharging(playerId, points, tx)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = l.ds.UpdatePlayerBalance(playerId, points, true, tx)
	if err != nil {
		log.Println(err.Error())
		err = errors.New(ErrInternalError)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (l *Logic) Fund(playerId int64, points int) error {
	tx, err := l.ds.BeginTx()
	if err != nil {
		log.Println(err.Error())
	}

	defer tx.Rollback()

	_, err = l.checkPlayer(playerId, tx)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = l.ds.UpdatePlayerBalance(playerId, points, false, tx)
	if err != nil {
		log.Println(err.Error())
		err = errors.New(ErrInternalError)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (l *Logic) AnnounceTournament(name string, deposit int) (int64, error) {
	tx, err := l.ds.BeginTx()
	if err != nil {
		log.Println(err.Error())
	}

	defer tx.Rollback()

	t := &data.Tournament{
		Name:    name,
		Deposit: deposit,
	}

	tid, err := l.ds.CreateNewTournament(t, tx)
	if err != nil {
		log.Println(err.Error())
		err = errors.New(ErrInternalError)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}

	return tid, err
}

func (l *Logic) JoinTournament(tournamentId int64, playerId int64, backerIds *[]int64) error {
	tx, err := l.ds.BeginTx()
	if err != nil {
		log.Println(err.Error())
	}

	defer tx.Rollback()

	t, err := l.checkTournament(tournamentId, tx)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	p, err := l.checkPlayer(playerId, tx)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if len(*backerIds) == 0 {
		err = l.checkBalanceForBacking(playerId, tournamentId, tx)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		err = l.ds.UpdatePlayerBalance(playerId, t.Deposit, true, tx)
		if err != nil {
			log.Println(err.Error())
			return errors.New(ErrInternalError)
		}

	} else {

		err = l.checkPlayers(backerIds, tx)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		backerIds, err = l.checkBackersInTournament(tournamentId, backerIds, tx)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		sum := (int)(t.Deposit-p.Points) / len(*backerIds)
		backers := l.getBackers(backerIds, playerId, sum)

		err = l.ds.BackPlayerForTournament(backers, tx)
		if err != nil {
			log.Println(err.Error())
			return errors.New(ErrInternalError)
		}

		err = l.ds.UpdatePlayersBalance(backerIds, sum, tx)
		if err != nil {
			log.Println(err.Error())
			return errors.New(ErrInternalError)
		}

		err = l.ds.UpdatePlayerBalance(playerId, t.Deposit-p.Points, true, tx)
		if err != nil {
			log.Println(err.Error())
			return errors.New(ErrInternalError)
		}

	}

	tp := &data.TournamentPlayer{
		TournamentId: tournamentId,
		PlayerId:     playerId,
	}

	_, err = l.ds.AddPlayerToTournament(tp, tx)
	if err != nil {
		log.Println(err.Error())
		return errors.New(ErrInternalError)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (l *Logic) Balance(playerId int64) (*data.PlayerBalance, error) {
	tx, err := l.ds.BeginTx()
	if err != nil {
		log.Println(err.Error())
	}

	defer tx.Rollback()

	_, err = l.checkPlayer(playerId, tx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	balance, err := l.ds.GetBalanceForPlayer(playerId, tx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	pb := &data.PlayerBalance{
		PlayerId: playerId,
		Balance:  balance,
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}

	return pb, err
}

func (l *Logic) ResultTournament(tournamentId int64) (*data.TournamentResults, error) {
	tx, err := l.ds.BeginTx()
	if err != nil {
		log.Println(err.Error())
	}

	defer tx.Rollback()

	_, err = l.checkTournament(tournamentId, tx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	winners, err := l.ds.GetTournamentWinners(tournamentId, tx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	playerPrizes := make([]data.PlayerPrize, 0)

	for _, w := range *winners {
		pp := data.PlayerPrize{
			PlayerId: w.PlayerId,
			Prize:    w.Prize,
		}

		playerPrizes = append(playerPrizes, pp)
	}

	tr := &data.TournamentResults{
		TournamentId: tournamentId,
		Winners:      playerPrizes,
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}

	return tr, err
}

func (l *Logic) UpdatePrizes(tournamentId int64, playerId int64, prize int) error {
	tx, err := l.ds.BeginTx()
	if err != nil {
		log.Println(err.Error())
	}

	defer tx.Rollback()

	_, err = l.checkTournament(tournamentId, tx)
	if err != nil {
		log.Println(err.Error())
		return errors.New(ErrInternalError)
	}

	_, err = l.checkPlayer(playerId, tx)
	if err != nil {
		log.Println(err.Error())
		return errors.New(ErrInternalError)
	}

	tp := &data.TournamentPlayer{
		TournamentId: tournamentId,
		PlayerId:     playerId,
		Prize:        prize,
	}

	err = l.ds.SetPlayerPrize(tp, tx)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}
	return err

}

func (l *Logic) Reset() error {
	err := l.ds.ClearDB()
	if err != nil {
		log.Println(err.Error())
		return errors.New(ErrInternalError)
	}
	return err
}

func (l *Logic) checkPlayer(playerId int64, tx *sql.Tx) (*data.Player, error) {
	p, err := l.ds.GetPlayerById(playerId, tx)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(ErrInternalError)
	}

	if p.Id == 0 {
		return nil, errors.New(ErrPlayerNotFound)
	}

	return p, err
}

func (l *Logic) checkPlayerInTournament(playerId, tournamentId int64, tx *sql.Tx) error {
	pId, err := l.ds.GetTournamentPlayerIdFromTournament(playerId, tournamentId, tx)
	if err != nil {
		log.Println(err.Error())
		return errors.New(ErrInternalError)
	}

	if pId == 0 {
		return errors.New(ErrPlayerAlreadyInTournament)
	}

	return err
}

func (l *Logic) checkBalanceForCharging(playerId int64, chargePoints int, tx *sql.Tx) error {
	balance, err := l.ds.GetBalanceForPlayer(playerId, tx)
	if err != nil {
		log.Println(err.Error())
		return errors.New(ErrInternalError)
	}

	if balance < chargePoints {
		return errors.New(ErrInsufficientBalance)
	}

	return err
}

func (l *Logic) checkBalanceForBacking(playerId int64, tournamentId int64, tx *sql.Tx) error {
	balance, err := l.ds.GetBalanceForPlayer(playerId, tx)
	if err != nil {
		log.Println(err.Error())
		return errors.New(ErrInternalError)
	}

	t, err := l.ds.GetTournamentById(tournamentId, tx)
	if err != nil {
		log.Println(err.Error())
		return errors.New(ErrInternalError)
	}

	if balance < t.Deposit {
		return errors.New(ErrInsufficientBalance)
	}

	return err
}

func (l *Logic) checkTournament(tournamentId int64, tx *sql.Tx) (*data.Tournament, error) {
	t, err := l.ds.GetTournamentById(tournamentId, tx)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(ErrInternalError)
	}

	if t.Id == 0 {
		return nil, errors.New(ErrTournamentNotFound)
	}

	return t, err
}

func (l *Logic) checkPlayers(players *[]int64, tx *sql.Tx) error {
	b, err := l.ds.GetPlayersByIds(players, tx)
	if err != nil {
		log.Println(err.Error())
		return errors.New(ErrInternalError)
	}

	if len(*b) != len(*players) {
		log.Println(err.Error())
		return errors.New(ErrPlayerNotFound)
	}

	return err
}

func (l *Logic) checkBackersInTournament(tournamentId int64, backers *[]int64, tx *sql.Tx) (*[]int64, error) {
	players, err := l.ds.GetPlayersByTournament(tournamentId, tx)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(ErrInternalError)
	}

	backersMap := make(map[int64]bool)

	for _, player := range *players {
		if arrayContains(backers, player) {
			backersMap[player] = true
		} else {
			backersMap[player] = false
		}
	}

	// only existing ones
	cleanBackers := make([]int64, 0)
	for k, v := range backersMap {
		if v {
			cleanBackers = append(cleanBackers, k)
		}

	}

	return &cleanBackers, err
}

func arrayContains(s *[]int64, e int64) bool {
	for _, a := range *s {
		if a == e {
			return true
		}
	}
	return false
}

func (l *Logic) getBackers(backerIds *[]int64, playerId int64, sum int) *[]data.Backer {
	backers := make([]data.Backer, 0)

	for _, bid := range *backerIds {
		b := data.Backer{
			BackerId: bid,
			PlayerId: playerId,
			Sum:      sum,
		}
		backers = append(backers, b)
	}

	return &backers
}
