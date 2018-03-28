package data

// Player of the service
//
// swagger:model Player
type Player struct {
	// Player's id.
	Id int64 `json:"id"`

	// Player's name.
	Name string `json:"name"`

	// Player's points.
	Points int `json:"points"`
}

// Tournament on the service
//
// swagger:model Tournament
type Tournament struct {

	// Tournaments's id.
	Id int64 `json:"id"`

	// Tournaments's name.
	Name string `json:"name"`

	// Tournaments's deposit.
	Deposit int `json:"deposit"`
}

// Tournament and its players on the service
//
// swagger:model TournamentPlayer
type TournamentPlayer struct {

	// TP id.
	Id           int64 `json:"id"`

	// Tournaments's id.
	TournamentId int64 `json:"tournament_id"`

	// The id of the player.
	PlayerId     int64 `json:"player_id"`

	// Prize of the player.
	Prize        int   `json:"prize"`
}

// Results of the tournament
//
// swagger:model TournamentResults
type TournamentResults struct {
	// The id of the tournament.
	//
	// required: true
	//
	TournamentId int64 `json:"tournament_id"`

	// The list Of winners.
	Winners []PlayerPrize `json:"winners"`
}

// Prize of the player
//
// swagger:model PlayerPrize
type PlayerPrize struct {
	// The id of the player.
	PlayerId int64 `json:"player_id"`

	// Value of the prize.
	Prize int `json:"prize"`
}

// Balance of the player
//
// swagger:model PlayerBalance
type PlayerBalance struct {
	// The id of the player
	PlayerId int64 `json:"player_id"`

	// The balance of the player.
	Balance int `json:"balance"`
}

type Backer struct {
	Id       int64 `json:"id"`
	PlayerId int64 `json:"player_id"`
	BackerId int64 `json:"backer_id"`
	Sum      int   `json:"sum"`
}
