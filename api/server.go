//     Schemes: http, https
//     Host: localhost:8080
//     BasePath: /v2
//     Version: 1.0.0
//     Title: SocialTournamentService API
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Alex Bondar<abondar1992@gmail.com>
//
//     Produces:
//     - application/json
//
// swagger:meta
package api

import (
	"encoding/json"
	"github.com/abondar24/SocialTournamentService/blogic"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"github.com/abondar24/SocialTournamentService/data"
)

type Server struct {
	logic  *blogic.Logic
	router *mux.Router
}

const (
	ErrPlayerNotFound            = "player not found"
	ErrTournamentNotFound        = "tournament not found"
	ErrInternalError             = "internal error"
	ErrInsufficientBalance       = "not enough points"
	ErrPlayerAlreadyInTournament = "player is already participating participating in tournament"
)

func NewServer(logic *blogic.Logic) *Server {
	router := mux.NewRouter().StrictSlash(true)
	return &Server{
		logic,
		router,
	}
}

func (s *Server) RunRestServer() {
	s.router.HandleFunc("/v2/", s.Index)
	s.router.HandleFunc("/v2/add_player", s.AddPlayer).Methods("POST")
	s.router.HandleFunc("/v2/take", s.Take).Methods("PUT")
	s.router.HandleFunc("/v2/fund", s.Fund).Methods("PUT")
	s.router.HandleFunc("/v2/announce_tournament", s.AnnounceTournament).Methods("POST")
	s.router.HandleFunc("/v2/join_tournament", s.JoinTournament).Methods("PUT")
	s.router.HandleFunc("/v2/result_tournament/{tId}", s.ResultTournament).Methods("GET")
	s.router.HandleFunc("/v2/balance/{pId}", s.Balance).Methods("GET")
	s.router.HandleFunc("/v2/reset", s.Reset).Methods("GET")
	s.router.HandleFunc("/v2/update_prizes", s.UpdatePrizes).Methods("PUT")
	s.router.HandleFunc("/v2/get_players", s.GetPlayers).Methods("GET")
	s.router.HandleFunc("/v2/get_tournaments", s.GetTournaments).Methods("GET")
	s.router.HandleFunc("/v2/get_players_tournament/{tId}", s.GetPlayersInTournament).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", s.router))

}

// Index swagger:route GET / Index
//
// Test server is up.
//
// Responses:
//    200: rsUp
func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is running")
}

// AddPlayer swagger:route POST /add_player  AddPlayer
//
// Add a new player.
//
// Responses:
//    201: rsCreated
//    500: errInternalError
func (s *Server) AddPlayer(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query()["name"][0]
	points := r.URL.Query()["points"][0]

	pts, err := strconv.Atoi(points)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	pId, err := s.logic.AddPlayer(name, pts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(pId)
}

// Take swagger:route PUT /take Take
//
// Charge a player player.
//
// Responses:
//    200: rsBalanceChanged
//    404: errPlayerTournamentNotFound
//    402: errInsufficientBalance
//    500: errInternalError
func (s *Server) Take(w http.ResponseWriter, r *http.Request) {
	playerId := r.URL.Query()["player_id"][0]
	points := r.URL.Query()["points"][0]

	pid, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	pts, err := strconv.Atoi(points)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	err = s.logic.Take(pid, pts)
	if err != nil {
		if err.Error() == ErrPlayerNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrInsufficientBalance {
			w.WriteHeader(http.StatusPaymentRequired)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	w.WriteHeader(http.StatusOK)
}

// Fund swagger:route PUT /fund Fund
//
// Add points to player.
//
// Responses:
//    200: rsBalanceChanged
//    404: errPlayerTournamentNotFound
//    500: errInternalError
func (s *Server) Fund(w http.ResponseWriter, r *http.Request) {
	playerId := r.URL.Query()["player_id"][0]
	points := r.URL.Query()["points"][0]

	pid, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	pts, err := strconv.Atoi(points)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	err = s.logic.Fund(pid, pts)
	if err != nil {
		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrPlayerNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}
	}

	w.WriteHeader(http.StatusOK)
}

// AnnounceTournament swagger:route POST /announce_tournament AnnounceTournament
//
// Announce a new tournament.
//
// Responses:
//    201: rsCreated
//    404: errPlayerTournamentNotFound
//    500: errInternalError
func (s *Server) AnnounceTournament(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()["name"][0]
	deposit := r.URL.Query()["deposit"][0]

	dp, err := strconv.Atoi(deposit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	tId, err := s.logic.AnnounceTournament(name, dp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tId)
}

// JoinTournament swagger:route PUT /join_tournament JoinTournament
//
// Join tournament with backers.
//
// Responses:
//    200: rsPlayerInTournament
//    402: errInsufficientBalance
//    409: errPlayerAlreadyInTournament
//    404: errPlayerTournamentNotFound
//    500: errInternalError
func (s *Server) JoinTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId := r.URL.Query()["tournament_id"][0]
	playerId := r.URL.Query()["player_id"][0]
	backerIds := r.URL.Query()["backer_id"]

	tid, err := strconv.ParseInt(tournamentId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	pid, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	backers, err := s.convertToInt64(backerIds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	err = s.logic.JoinTournament(tid, pid, backers)
	if err != nil {
		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrTournamentNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrPlayerNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrInsufficientBalance {
			w.WriteHeader(http.StatusPaymentRequired)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrPlayerAlreadyInTournament {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
		}
	}

	w.WriteHeader(http.StatusOK)
}

// UpdatePrizes swagger:route PUT /update_prizes
//
// Updates player's prize by a getting a list of TournamentPlayer objects
//
// Responses:
//    200: rsPlayerInTournament
//    500: errInternalError
func (s *Server) UpdatePrizes(w http.ResponseWriter, r *http.Request) {


	decoder := json.NewDecoder(r.Body)
	var tps []data.TournamentPlayer

	err := decoder.Decode(&tps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer r.Body.Close()

	err = s.logic.UpdatePrizes(&tps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)

}

// ResultTournament swagger:route GET /result_tournament/{tId} ResultTournament
//
// Get results of tournament.
//
// Responses:
//    200: rsResultTournament
//    404: errPlayerTournamentNotFound
//    500: errInternalError
func (s *Server) ResultTournament(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentId := vars["tId"]

	tid, err := strconv.ParseInt(tournamentId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	tr, err := s.logic.ResultTournament(tid)
	if err != nil {
		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrTournamentNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}
	}

	json.NewEncoder(w).Encode(tr)
}

// Balance swagger:route GET /balance/{pId}  Balance
//
// Returns player's balance
//
// Responses:
//    200: rsBalance
//    404: errPlayerTournamentNotFound
//    500: errInternalError
func (s *Server) Balance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerId := vars["pId"]

	pid, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	pb, err := s.logic.Balance(pid)
	if err != nil {
		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrPlayerNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}
	}

	json.NewEncoder(w).Encode(pb)
}

// Rest swagger:route GET /reset Reset
// Resets database
//
// Responses:
//    200: rsDbReset
//    500: errInternalError
func (s *Server) Reset(w http.ResponseWriter, r *http.Request) {
	err := s.logic.Reset()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}

// AddPlayer swagger:route GET /get_players  GetPlayers
//
// Get all players.
//
// Responses:
//    200: rsPlayers
//    500: errInternalError
func (s *Server) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := s.logic.GetAllPlayers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(players)
}

// GetTournaments swagger:route GET /get_tournaments  GetTournaments
//
// Get all tournaments.
//
// Responses:
//    200: rsTournaments
//    500: errInternalError
func (s *Server) GetTournaments(w http.ResponseWriter, r *http.Request) {
	tournaments, err := s.logic.GetAllTournaments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tournaments)
}

// GetPlayersInTournament swagger:route GET /get_players_tournament/{tId} GetPlayersInTournament
//
// Get players taking part in tournament.
//
// Responses:
//    200: rsPlayers
//    404: errPlayerTournamentNotFound
//    500: errInternalError
func (s *Server) GetPlayersInTournament(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentId := vars["tId"]

	tid, err := strconv.ParseInt(tournamentId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	players, err := s.logic.GetPlayersTournament(tid)
	if err != nil {
		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		if err.Error() == ErrTournamentNotFound {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusNotFound)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

func (s *Server) convertToInt64(strArray []string) (*[]int64, error) {
	int64Array := make([]int64, 0)
	for _, backerId := range strArray {
		backer, err := strconv.ParseInt(backerId, 10, 64)
		if err != nil {
			return nil, err
		}
		int64Array = append(int64Array, backer)
	}

	return &int64Array, nil
}
