package api

import "github.com/abondar24/SocialTournamentService/data"

// Resource not found
// swagger:response errPlayerTournamentNotFound
type ErrorPlayerOrTournamentNotFound struct {
	// The error message
	// in: body
	Body struct {
		// "Player or Tournament not found in database"
		//
		// Required: true
		Message string
	}
}

// Internal error
// swagger:response errInternalError
type ErrorInternalError struct {
	// The error message
	// in: body
	Body struct {
		// "Internal error on the server"
		//
		// Required: true
		Message string
	}
}

// Not enough points
// swagger:response errInsufficientBalance
type ErrorInsufficientBalance struct {
	// The error message
	// in: body
	Body struct {
		// "Not enough points to join tournament"
		//
		// Required: true
		Message string
	}
}

// Player is already participating in tournament
// swagger:response errPlayerAlreadyInTournament
type ErrorPlayerAlreadyInTournament struct {
	// The error message
	// in: body
	Body struct {
		// "Player is already participating in tournament"
		//
		// Required: true
		Message string
	}
}

// Server up
// swagger:response rsUp
type ResponseServerUp struct {
	// The error message
	// in: body
	Body struct {
		// "Server is running"
		//
		// Required: true
		Message string
	}
}

// Resource created
// swagger:response rsCreated
type ResponseResourceCreated struct {
	// The error message
	// in: body
	Body struct {
		// "Player or tournament has been added"
		//
		// Required: true
		Message string
	}
}

// Balance changed
// swagger:response rsBalanceChanged
type ResponseBalanceChanged struct {
	// The error message
	// in: body
	Body struct {
		// "Player's points added or removed"
		//
		// Required: true
		Message string
	}
}

// Player added to tournament
// swagger:response rsPlayerInTournament
type ResponsePlayerInTournament struct {
	// The error message
	// in: body
	Body struct {
		// "Player has been added to tournament"
		//
		// Required: true
		Message string
	}
}

// Prize updated
// swagger:response rsPlayerInTournament
type ResponsePrizeUpdated struct {
	// The error message
	// in: body
	Body struct {
		// "Player's prize has been updated"
		//
		// Required: true
		Message string
	}
}

// Tournament results
// swagger:response rsResultTournament
type ResponseResultTournament struct {
	// in: body
	Data data.TournamentResults `json:"data"`
}

// Player's balance results
// swagger:response rsBalance
type ResponseBalance struct {
	// in: body
	Data data.PlayerBalance `json:"data"`
}

// Players in system
// swagger:response rsPlayers
type ResponsePlayers struct {
	// in: body
	Data []data.Player `json:"data"`
}

// Tournaments in system
// swagger:response rsTournaments
type ResponseTournaments struct {
	// in: body
	Data []data.Tournament `json:"data"`
}

// Database is reset
// swagger:response rsDbReset
type ResponseDatabaseReset struct {
	// The error message
	// in: body
	Body struct {
		// "database is reset successfully"
		//
		// Required: true
		Message string
	}
}
