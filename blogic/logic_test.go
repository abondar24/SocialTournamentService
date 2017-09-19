package blogic

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/abondar24/SocialTournamentService/data"
)

func Test_Logic_Join_Tournament(t *testing.T) {

	ds, err := data.ConnectToBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	p := &data.Player{
		Name:   "Ahmed",
		Points: 100,
	}


	to := &data.Tournament{
		Name:    "blacjack",
		Deposit: 50,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}


	tId, err := ds.CreateNewTournament(to)
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{}

	err = l.JoinTournament(tId,pId,backerIds )

	assert.Equal(t, nil,err)
}



func Test_Logic_Join_Tournament_InsufficientFunds(t *testing.T) {

	ds, err := data.ConnectToBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	p := &data.Player{
		Name:   "Ahmed",
		Points: 100,
	}


	to := &data.Tournament{
		Name:    "blacjack",
		Deposit: 200,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}


	tId, err := ds.CreateNewTournament(to)
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{}

	err = l.JoinTournament(tId,pId,backerIds )

	assert.Equal(t, "not enough points",err.Error())
}

func Test_Logic_Join_Tournament_With_Backers(t *testing.T) {

	ds, err := data.ConnectToBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	p := &data.Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pb := &data.Player{
		Name:   "Rudi",
		Points: 1000,
	}

	pb1 := &data.Player{
		Name:   "Hans",
		Points: 1000,
	}

	to := &data.Tournament{
		Name:    "blacjack",
		Deposit: 200,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	pbId, err := ds.CreateNewPlayer(pb)
	if err != nil {
		t.Fatal(err)
	}

	pb1Id, err := ds.CreateNewPlayer(pb1)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := ds.CreateNewTournament(to)
	if err != nil {
		t.Fatal(err)
	}

	tp := &data.TournamentPlayer{
		PlayerId:     pbId,
		TournamentId: tId,
	}

	tp1 := &data.TournamentPlayer{
		PlayerId:     pb1Id,
		TournamentId: tId,
	}

	_, err = ds.AddPlayerToTournament(tp)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ds.AddPlayerToTournament(tp1)
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{pbId, pb1Id}

	err = l.JoinTournament(tId,pId,backerIds)

	assert.Equal(t, nil,err)
}


func Test_Logic_Join_Tournament_With_Backers_Not_In_Tournament(t *testing.T) {

	ds, err := data.ConnectToBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	p := &data.Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pb := &data.Player{
		Name:   "Rudi",
		Points: 1000,
	}

	pb1 := &data.Player{
		Name:   "Hans",
		Points: 1000,
	}

	to := &data.Tournament{
		Name:    "blacjack",
		Deposit: 200,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	pbId, err := ds.CreateNewPlayer(pb)
	if err != nil {
		t.Fatal(err)
	}

	pb1Id, err := ds.CreateNewPlayer(pb1)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := ds.CreateNewTournament(to)
	if err != nil {
		t.Fatal(err)
	}

	tp := &data.TournamentPlayer{
		PlayerId:     pbId,
		TournamentId: tId,
	}


	_, err = ds.AddPlayerToTournament(tp)
	if err != nil {
		t.Fatal(err)
	}



	backerIds := &[]int64{pbId, pb1Id}

	err = l.JoinTournament(tId,pId,backerIds )

	assert.Equal(t, "backer is not participating in tournament",err.Error())
}


func Test_Get_Tournament_Results(t *testing.T) {

	ds, err := data.ConnectToBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	l := NewLogic(ds)

	p := &data.Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pb := &data.Player{
		Name:   "Rudi",
		Points: 1000,
	}



	to := &data.Tournament{
		Name:    "blacjack",
		Deposit: 200,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	pbId, err := ds.CreateNewPlayer(pb)
	if err != nil {
		t.Fatal(err)
	}


	tId, err := ds.CreateNewTournament(to)
	if err != nil {
		t.Fatal(err)
	}

	backerIds := &[]int64{pbId}

	err = l.JoinTournament(tId,pbId,&[]int64{})
	if err != nil {
		t.Fatal(err)
	}

	err = l.JoinTournament(tId,pId,backerIds)
	if err != nil {
		t.Fatal(err)
	}

	err = l.UpdatePrizes(tId,pId,50)
	if err != nil {
		t.Fatal(err)
	}

	err = l.UpdatePrizes(tId,pbId,50)
	if err != nil {
		t.Fatal(err)
	}

	tr,err :=l.ResultTournament(tId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tId,tr.TournamentId)
	assert.Equal(t, 2,len(tr.Winners))
}


func Test_Clear_DataBase(t *testing.T) {
	ds, err := data.ConnectToBase()
	if err != nil {
		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}
	err = ds.ClearDB()


}