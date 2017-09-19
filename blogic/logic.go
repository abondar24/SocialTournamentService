package blogic

import (
	"errors"
	"github.com/abondar24/SocialTournamentService/data"
)

const (
	ErrPlayerNotFound          string = "player not found"
	ErrTournamentNotFound      string = "tournament not found"
	ErrInternalError           string = "internal error"
	ErrInsufficientBalance     string = "not enough points"
	ErrBackerIsNotInTournament string = "backer is not participating in tournament"
)

type Logic struct {
	ds *data.MySql
}

func NewLogic(dataSource *data.MySql) *Logic {

	return &Logic{
		dataSource,
	}
}

func (l *Logic) AddPlayer(name string, points int) error {
	p := &data.Player{
		Name:   name,
		Points: points,
	}

	_, err := l.ds.CreateNewPlayer(p)

	return err
}

func (l *Logic) Take(playerId int64, points int) error {
	_, err := l.checkPlayer(playerId)
	if err != nil {
		return err
	}

	err = l.checkBalanceForCharging(playerId, points)
	if err != nil {
		return err
	}

	err = l.ds.UpdatePlayerBalance(playerId, points, true)
	if err != nil {
		err = errors.New(ErrInternalError)
	}


	return err
}

func (l *Logic) Fund(playerId int64, points int) error {
	_, err := l.checkPlayer(playerId)
	if err != nil {
		return err
	}

	err = l.ds.UpdatePlayerBalance(playerId, points, false)
	if err != nil {
		err = errors.New(ErrInternalError)
	}

	return err
}

func (l *Logic) AnnounceTournament(name string, deposit int) error {
	t := &data.Tournament{
		Name:    name,
		Deposit: deposit,
	}

	_, err := l.ds.CreateNewTournament(t)
	if err != nil {
		err = errors.New(ErrInternalError)
	}

	return err
}

func (l *Logic) JoinTournament(tournamentId int64, playerId int64, backerIds *[]int64) error {
	t, err := l.checkTournament(tournamentId)
	if err != nil {
		return err
	}

	p, err := l.checkPlayer(playerId)
	if err != nil {
		return err
	}

	if len(*backerIds) == 0 {
		err = l.checkBalanceForBacking(playerId, tournamentId)
		if err != nil {
			return err
		}
	} else {

		err = l.checkPlayers(backerIds)
		if err != nil {
			return err
		}

		backerIds, err = l.checkBackersAreInTournament(tournamentId, backerIds)
		if err != nil {
			return err
		}

		sum := (int)(t.Deposit-p.Points) / len(*backerIds)
		backers := l.getBackers(backerIds, playerId, sum)

		err = l.ds.BackPlayerForTournament(backers)
		if err != nil {
			return errors.New(ErrInternalError)
		}

	}

	tp := &data.TournamentPlayer{
		TournamentId: tournamentId,
		PlayerId:     playerId,
	}

	_, err = l.ds.AddPlayerToTournament(tp)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	return err
}

func (l *Logic) Balance(playerId int64) (*data.PlayerBalance, error) {
	_, err := l.checkPlayer(playerId)
	if err != nil {
		return nil, err
	}

	balance, err := l.ds.GetBalanceForPlayer(playerId)
	if err != nil {
		return nil, err
	}

	pb := &data.PlayerBalance{
		PlayerId: playerId,
		Balance:  balance,
	}

	return pb, err
}

func (l *Logic) ResultTournament(tournamentId int64) (*data.TournamentResults, error) {
	_, err := l.checkTournament(tournamentId)
	if err != nil {
		return nil, err
	}

	winners, err := l.ds.GetTournamentWinners(tournamentId)
	if err != nil {
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

	return tr, err
}

func (l *Logic) Reset() error {
	err := l.ds.ClearDB()

	return err
}

func (l *Logic) UpdatePrizes(tournamentId int64,playerId int64,prize int) error {
	_,err:=l.checkTournament(tournamentId)
	if err != nil {
		return errors.New(ErrInternalError)
	}


	_,err = l.checkPlayer(playerId)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	tp:= &data.TournamentPlayer{
		TournamentId:tournamentId,
		PlayerId:playerId,
		Prize:prize,
	}

	err = l.ds.SetPlayerPrize(tp)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	return err

}

func (l *Logic) checkPlayer(playerId int64) (*data.Player, error) {
	p, err := l.ds.GetPlayerById(playerId)
	if err != nil {
		return nil, errors.New(ErrInternalError)
	}

	if p.Id == 0 {
		return nil, errors.New(ErrPlayerNotFound)
	}

	return p, err
}

func (l *Logic) checkBalanceForCharging(playerId int64, chargePoints int) error {
	balance, err := l.ds.GetBalanceForPlayer(playerId)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	if balance < chargePoints {
		return errors.New(ErrInsufficientBalance)
	}

	return err
}

func (l *Logic) checkBalanceForBacking(playerId int64, tournamentId int64) error {
	balance, err := l.ds.GetBalanceForPlayer(playerId)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	t, err := l.ds.GetTournamentById(tournamentId)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	if balance < t.Deposit {
		return errors.New(ErrInsufficientBalance)
	}

	return err
}

func (l *Logic) checkTournament(tournamentId int64) (*data.Tournament, error) {
	t, err := l.ds.GetTournamentById(tournamentId)
	if err != nil {
		return nil, errors.New(ErrInternalError)
	}

	if t.Id == 0 {
		return nil, errors.New(ErrTournamentNotFound)
	}

	return t, err
}

func (l *Logic) checkPlayers(players *[]int64) error {
	b, err := l.ds.GetPlayersByIds(players)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	if len(*b) != len(*players) {
		return errors.New(ErrPlayerNotFound)
	}

	return err
}

func (l *Logic) checkBackersAreInTournament(tournamentId int64, backers *[]int64) (*[]int64, error) {
	players, err := l.ds.GetTournamentPlayersIdsByTournamentId(tournamentId)
	if err != nil {
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
	newBackers := make([]int64, 0)
	for k, v := range backersMap {
		if v {
			newBackers = append(newBackers, k)
		}

	}

	if len(newBackers) != len(*backers) {
		err = errors.New(ErrBackerIsNotInTournament)
	}

	return &newBackers, err
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
			BackerId:bid,
			PlayerId: playerId,
			Sum:      sum,
		}
		backers = append(backers, b)
	}

	return &backers
}
