package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/abondar24/SocialTournamentService/data"
	"strconv"
	"errors"
)

type Server struct {
	ds     *data.MySql
	router *mux.Router
}

const (
	ErrPlayerNotFound      string = "player not found"
	ErrTournamentNotFound  string = "tournament not found"
	ErrInternalError       string = "internal error"
	ErrInsufficientBalance string = "not enough points"
)

func NewServer(dataSource *data.MySql) *Server {
	router := mux.NewRouter()

	return &Server{
		dataSource,
		router,
	}
}

func (s *Server) RunRestServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", s.Index)
	router.HandleFunc("/addPlayer/{name}/{points}", s.AddPlayer).Methods("GET")
	router.HandleFunc("/take/{player_id}", s.Take).Methods("GET")
	router.HandleFunc("/fund/{player_id}/{points}", s.Fund).Methods("GET")
	router.HandleFunc("/announceTournament/{tournament_id}/{deposit}",
		s.AnnounceTournament).Methods("GET")
	router.HandleFunc("/joinTournament/{tournament_id}/{player_id}/", s.JoinTournament).Methods("GET")
	router.HandleFunc("/resultTournament/{tournament_id}/", s.ResultTournament).Methods("POST")
	router.HandleFunc("/balance/{player_id}", s.Balance).Methods("GET")
	router.HandleFunc("/reset/", s.Reset).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is running")
}

func (s *Server) AddPlayer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}


func (s *Server) Take(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player_id := vars["player_id"]
	points := vars["points"]

	pid, err := strconv.ParseInt(player_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.checkPlayer(pid)
	if err.Error() == ErrInternalError {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err.Error() == ErrPlayerNotFound {
		w.WriteHeader(http.StatusNotFound)
	}

	pts, err := strconv.Atoi(points)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.checkBalance(pid, pts)

	if err.Error() == ErrInsufficientBalance {
		w.WriteHeader(http.StatusPaymentRequired)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.ds.UpdatePlayerBalance(pid, pts, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Fund(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player_id := vars["player_id"]
	points := vars["points"]

	pid, err := strconv.ParseInt(player_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.checkPlayer(pid)
	if err.Error() == ErrInternalError {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err.Error() == ErrPlayerNotFound {
		w.WriteHeader(http.StatusNotFound)
	}

	pts, err := strconv.Atoi(points)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.ds.UpdatePlayerBalance(pid, pts, false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) AnnounceTournament(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (s *Server) JoinTournament(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (s *Server) ResultTournament(w http.ResponseWriter, r *http.Request) {


}

func (s *Server) Balance(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Reset(w http.ResponseWriter, r *http.Request) {
	err := s.ds.ClearDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) checkPlayer(playerId int64) error {
	p, err := s.ds.GetPlayerById(playerId)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	if p.Id == 0 {
		return errors.New(ErrPlayerNotFound)
	}

	return err
}

func (s *Server) checkBalance(playerId int64, chargePoints int) error {
	p, err := s.ds.GetPlayerById(playerId)
	if err != nil {
		return errors.New(ErrInternalError)
	}

	if p.Id == 0 {
		return errors.New(ErrPlayerNotFound)
	}

	return err
}
