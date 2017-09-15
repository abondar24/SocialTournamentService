package data


type Player struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Balance int `json:"balance"`
	BackId int64 `json:"back_id"`
}

type Tournament struct {
	Id int64 `json:"id"`
	TournamentName string `json:"tournament_name"`
	Deposit int `json:"deposit"`
}

type TournamentPlayer struct {
	Id int64 `json:"id"`
	TournamentId int64 `json:"tournament_id"`
	PlayerId int64 `json:"player_id"`
	Prize int `json:"prize"`
}

type Backer struct {
	Id int64 `json:"id"`
	PlayerId int64 `json:"player_id"`
	BackerId int64 `json:"backer_id"`
	Sum int `json:"sum"`
}



