package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Clear_DataBase(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {

		t.Fatal(err)
	}

	assert.Equal(t, nil, err)

}

func Test_Get_Create_New_Player(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	p := &Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, pId > 0)

	pl, err := ds.GetPlayerById(pId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, p.Name, pl.Name)
	assert.Equal(t, p.Points, pl.Points)

}


func Test_Get_Players_By_Id(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	p := &Player{
		Name:   "Ahmed",
		Points: 100,
	}

	p1 := &Player{
		Name:   "Rudi",
		Points: 100,
	}


	p2 := &Player{
		Name:   "Mark",
		Points: 100,
	}

	playerIds := make([]int64,0)

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	playerIds = append(playerIds,pId)

	p1Id, err := ds.CreateNewPlayer(p1)
	if err != nil {
		t.Fatal(err)
	}

	playerIds = append(playerIds,p1Id)


	p2Id, err := ds.CreateNewPlayer(p2)
	if err != nil {
		t.Fatal(err)
	}

	playerIds = append(playerIds,p2Id)

	pls, err := ds.GetPlayersByIds(&playerIds)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3,len(*pls))

}

func Test_Get_Player_Balance(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	p := &Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	balance, err := ds.GetBalanceForPlayer(pId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, p.Points, balance)

}

func Test_Add_To_Balance(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	p := &Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	err = ds.UpdatePlayerBalance(pId, 500, false)
	if err != nil {
		t.Fatal(err)
	}

	balance, err := ds.GetBalanceForPlayer(pId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 600, balance)
}

func Test_Take_From_Balance(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	p := &Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	err = ds.UpdatePlayerBalance(pId, 50, true)
	if err != nil {
		t.Fatal(err)
	}

	balance, err := ds.GetBalanceForPlayer(pId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 50, balance)
}

func Test_Add_New_Tournament(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	toExpected := &Tournament{
		Name:    "blacjack",
		Deposit: 20,
	}

	tId, err := ds.CreateNewTournament(toExpected)
	if err != nil {
		t.Fatal(err)
	}

	toActual, err := ds.GetTournamentById(tId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, toExpected.Name, toActual.Name)
	assert.Equal(t, toExpected.Deposit, toActual.Deposit)

}

func Test_Add_Player_To_Tournament(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	to := &Tournament{
		Name:    "blacjack",
		Deposit: 20,
	}

	p := &Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := ds.CreateNewTournament(to)
	if err != nil {
		t.Fatal(err)
	}

	tp := &TournamentPlayer{
		PlayerId:     pId,
		TournamentId: tId,
		Prize:        0,
	}

	tpId, err := ds.AddPlayerToTournament(tp)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, tpId > 0)

}

func Test_Get_Tournament_Players_By_TournamentId(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	to := &Tournament{
		Name:    "blacjack",
		Deposit: 20,
	}

	p := &Player{
		Name:   "Ahmed",
		Points: 100,
	}

	p1 := &Player{
		Name:   "Abdi",
		Points: 100,
	}

	p2 := &Player{
		Name:   "Rudolf",
		Points: 100,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	p1Id, err := ds.CreateNewPlayer(p1)
	if err != nil {
		t.Fatal(err)
	}

	p2Id, err := ds.CreateNewPlayer(p2)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := ds.CreateNewTournament(to)
	if err != nil {
		t.Fatal(err)
	}

	tp := &TournamentPlayer{
		PlayerId:     pId,
		TournamentId: tId,
		Prize:        0,
	}

	tp1 := &TournamentPlayer{
		PlayerId:     p1Id,
		TournamentId: tId,
		Prize:        0,
	}

	tp2 := &TournamentPlayer{
		PlayerId:     p2Id,
		TournamentId: tId,
		Prize:        0,
	}

	_, err = ds.AddPlayerToTournament(tp)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ds.AddPlayerToTournament(tp1)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ds.AddPlayerToTournament(tp2)
	if err != nil {
		t.Fatal(err)
	}

	players, err := ds.GetTournamentPlayersIdsByTournamentId(tId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(*players))

}

func Test_Set_Player_Prize(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	to := &Tournament{
		Name:    "blacjack",
		Deposit: 20,
	}

	p := &Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	tId, err := ds.CreateNewTournament(to)
	if err != nil {
		t.Fatal(err)
	}

	tp := &TournamentPlayer{
		PlayerId:     pId,
		TournamentId: tId,
		Prize:        0,
	}

	_, err = ds.AddPlayerToTournament(tp)
	if err != nil {
		t.Fatal(err)
	}

	tp.Prize = 500
	err = ds.SetPlayersPrize(tp)
	if err != nil {
		t.Fatal(err)
	}

	winners, err := ds.GetTournamentWinners(tId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(*winners))

	for _,w:= range *winners{
		assert.Equal(t, 500, w.Prize)
	}

}

func Test_Back_Player_Tournament(t *testing.T) {
	ds, err := ConnectToBase()
	if err != nil {

		t.Fatal(err)
	}

	err = ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	p := &Player{
		Name:   "Ahmed",
		Points: 100,
	}

	pb := &Player{
		Name:   "Rudi",
		Points: 1000,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	pbId, err := ds.CreateNewPlayer(pb)
	if err != nil {
		t.Fatal(err)
	}

	back := &Backer{
		PlayerId: pId,
		BackerId: pbId,
		Sum:      200,
	}

	backers:=[]Backer{*back}

	err = ds.BackPlayerForTournament(&backers)
	if err != nil {
		t.Fatal(err)
	}


	backers, err = ds.GetPlayerBackers(pId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(backers))
}
