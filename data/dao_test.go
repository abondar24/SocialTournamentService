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
		Name:    "Ahmed",
		Balance: 100,
		BackId:  0,
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
	assert.Equal(t, p.Balance, pl.Balance)

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
		Name:    "Ahmed",
		Balance: 100,
	}

	pId, err := ds.CreateNewPlayer(p)
	if err != nil {
		t.Fatal(err)
	}

	balance, err := ds.GetBalanceForPlayer(pId)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, p.Balance, balance)

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
		Name:    "Ahmed",
		Balance: 100,
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
		Name:    "Ahmed",
		Balance: 100,
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
