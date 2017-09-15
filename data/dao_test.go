package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMigrate_Clear_DataBase(t *testing.T) {
	ds := ConnectToBase()
	err := ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, nil, err)

}

func TestMigrate_Get_Player_By_Id(t *testing.T) {
	ds := ConnectToBase()
	err := ds.ClearDB()
	if err != nil {
		t.Fatal(err)
	}



	pId := int64(1)

	expPlayer := &Player{
		Id:      0,
		Name:    "",
		Balance: 0,
		BackId:  0,
	}

	p, err := ds.GetPlayerById(pId)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expPlayer, p)

}

func TestMigrate_Get_Create_New_Player(t *testing.T) {
	ds := ConnectToBase()
	err := ds.ClearDB()
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

	assert.True(t,  pId>0)

}
