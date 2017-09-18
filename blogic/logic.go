package blogic

import (
	"errors"
	"github.com/abondar24/SocialTournamentService/data"
	"fmt"
)


const (
	ErrPlayerNotFound         string = "player not found"
	ErrTournamentNotFound      string ="tournament not found"
	ErrInternalError           string = "internal error"
	ErrInsufficientBalance     string ="not enough points"
	ErrBackerIsNotInTournament string = "backer is not participating in tournament"
)

type Logic struct {
	ds     *data.MySql

}


func NewLogic(dataSource *data.MySql) *Logic {

	return &Logic{
		dataSource,
	}
}



func (l *Logic) AddPlayer(name string,points int) error{
	p := &data.Player{
		Name:   name,
		Points: points,
	}

	_,err := l.ds.CreateNewPlayer(p)

	return err
}


func (l *Logic) Take(playerId int64, points int) error {
	err := l.checkPlayer(playerId)
	fmt.Println(err)
	if err!=nil{
		return err
	}

	err = l.checkBalanceForCharging(playerId,points)
	if err!=nil{
		return err
	}

	err = l.ds.UpdatePlayerBalance(playerId, points, true)
	if err!=nil{
		err = errors.New(ErrInternalError)
	}

	fmt.Println(err)

	return err
}

func (l *Logic) Fund(playerId int64, points int) error {
	err := l.checkPlayer(playerId)
	if err!=nil{
		return err
	}


	err = l.ds.UpdatePlayerBalance(playerId, points, false)
	if err!=nil{
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
	if err!=nil{
		err = errors.New(ErrInternalError)
	}

	return err
}

func (l *Logic) JoinTournament(tournamentId int64, playerId int64, backers *[]int64) error {
	err := l.checkTournament(tournamentId)
	if err!=nil{
		return err
	}

	err = l.checkPlayer(playerId)
	if err!=nil{
		return err
	}

	err = l.checkBackers(backers)
	if err!=nil{
		return err
		}

	//if len(backers) == 0 {
	//	err = l.checkBalanceForBacking(playerId, tournamentId)
	//	if err != nil {
	//		return err
	//	}
	//} else {
	//
	//	err = l.backPlayer(pid, *backers)
	//
	//}

	return err
}



func (l *Logic) Balance(playerId int64) (*data.PlayerBalance,error){
	err := l.checkPlayer(playerId)
	if err!=nil{
		return nil,err
	}

	balance, err := l.ds.GetBalanceForPlayer(playerId)
	if err != nil {
	   return nil,err
	}

	pb := &data.PlayerBalance{
		PlayerId: playerId,
		Balance:  balance,
	}

	return pb,err
}

func (l *Logic) Reset() error{
	err := l.ds.ClearDB()

	return err
}


func (l *Logic) checkPlayer(playerId int64) error {
	p, err := l.ds.GetPlayerById(playerId)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	if p.Id == 0 {
		return errors.New(ErrPlayerNotFound)
	}

	return err
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

func (l *Logic) checkTournament(tournamentId int64) error {
	t, err := l.ds.GetTournamentById(tournamentId)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	if t.Id == 0 {
		return errors.New(ErrTournamentNotFound)
	}

	return err
}

func (l *Logic) checkBackers(backer *[]int64) error {
	b, err := l.ds.GetPlayersByIds(backer)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	if len(b)!=len(*backer) {
		return errors.New(ErrPlayerNotFound)
	}

	return err
}