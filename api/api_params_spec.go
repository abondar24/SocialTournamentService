package api

// swagger:parameters AddPlayer
type PlayerNameParams struct {
	// The name of the player.
	//
	// required: true
	Name string `json:"name"`

	// Points of the player.
	//
	// required: true
	Points int `json:"points"`
}

// swagger:parameters Take Fund
type PlayerIdParams struct {
	// Id of the player.
	//
	//
	Id int64 `json:"id"`

	// Points of the player.
	//
	// required: true
	Points int `json:"points"`
}

// swagger:parameters AnnounceTournament
type AnnounceTournamentParams struct {
	// The name of the tournament.
	//
	// required: true
	Name string `json:"name"`

	// Minimum number of points to take part in tournament.
	//
	// required: true
	Deposit int `json:"deposit"`
}

// swagger:parameters JoinTournament
type JoinTournamentParams struct {
	// The id of the tournament.
	//
	// required: true
	TournamentId int64 `json:"tournament_id"`

	// The id of the player.
	PlayerId int64 `json:"player_id"`

	// Id of players backer(can'be several backers)
	//
	// required: false
	BackerId int64 `json:"backer_id"`
}

// swagger:parameters UpdatePrize
type UpdatePrizeParams struct {
	// The id of the tournament.
	//
	// required: true
	TournamentId int64 `json:"tournament_id"`

	// The id of the player.
	PlayerId int64 `json:"player_id"`

	// Value of the prize.
	//
	// required: true
	Prize int `json:"prize"`
}

// swagger:parameters Balance
type PlayerIdParam struct {

	// Id of the player.
	//
	// required: true
	// in:path
	Id int64 `json:"pId"`
}

// swagger:parameters ResultTournament GetPlayersInTournament
type TournamentIdParam struct {

	// Id of the tournament.
	//
	// required: true
	// in:path
	Id int64 `json:"tId"`
}
