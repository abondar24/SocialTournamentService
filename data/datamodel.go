package data

type Player struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
	BackId int64  `json:"back_id"`
}

type Tournament struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Deposit int    `json:"deposit"`
}

type TournamentPlayer struct {
	Id           int64 `json:"id"`
	TournamentId int64 `json:"tournament_id"`
	PlayerId     int64 `json:"player_id"`
	Prize        int   `json:"prize"`
}

type TournamentResults struct {
	TournamentId int64         `json:"tournament_id"`
	Winners      []PlayerPrize `json:"winners"`
}

type PlayerPrize struct {
	PlayerId int64 `json:"player_id"`
	Prize    int   `json:"prize"`
}

type Backer struct {
	Id       int64 `json:"id"`
	PlayerId int64 `json:"player_id"`
	BackerId int64 `json:"backer_id"`
	Sum      int   `json:"sum"`
}
