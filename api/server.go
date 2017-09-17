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
	router := mux.NewRouter().StrictSlash(true)

	return &Server{
		dataSource,
		router,
	}
}

func (s *Server) RunRestServer() {


	s.router.HandleFunc("/", s.Index)
	s.router.HandleFunc("/add_player", s.AddPlayer).Methods("GET")
	s.router.HandleFunc("/take/", s.Take).Methods("GET")
	s.router.HandleFunc("/fund", s.Fund).Methods("GET")
	s.router.HandleFunc("/announce_tournament",
		s.AnnounceTournament).Methods("GET")
	s.router.HandleFunc("/join_tournament", s.JoinTournament).Methods("GET")
	s.router.HandleFunc("/result_tournament", s.ResultTournament).Methods("POST")
	s.router.HandleFunc("/balance", s.Balance).Methods("GET")
	s.router.HandleFunc("/reset", s.Reset).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", s.router))

}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is running")
}

func (s *Server) AddPlayer(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query()["name"][0]
	points :=  r.URL.Query()["points"][0]

	pts, err := strconv.Atoi(points)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	p:= &data.Player{
		Name: name,
		Points: pts,
	}


	_,err = s.ds.CreateNewPlayer(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}


func (s *Server) Take(w http.ResponseWriter, r *http.Request) {
	player_id := r.URL.Query()["player_id"][0]
	points :=  r.URL.Query()["points"][0]

	pid, err := strconv.ParseInt(player_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.checkPlayer(pid)
	if err!=nil{
		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err.Error() == ErrPlayerNotFound {
			w.WriteHeader(http.StatusNotFound)
		}

	}

	pts, err := strconv.Atoi(points)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.checkBalance(pid, pts)
	if err!=nil{
		if err.Error() == ErrInsufficientBalance {
			w.WriteHeader(http.StatusPaymentRequired)
		}

		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}


	err = s.ds.UpdatePlayerBalance(pid, pts, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Fund(w http.ResponseWriter, r *http.Request) {
	player_id := r.URL.Query()["player_id"][0]
	points :=  r.URL.Query()["points"][0]

	pid, err := strconv.ParseInt(player_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.checkPlayer(pid)
	if err!=nil{
		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err.Error() == ErrPlayerNotFound {
			w.WriteHeader(http.StatusNotFound)
		}
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

	w.WriteHeader(http.StatusOK)
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

	balance,err := s.ds.GetBalanceForPlayer(p.Id)
	if err != nil {
		return errors.New(ErrInternalError)
	}


	if balance<chargePoints{
		return errors.New(ErrInsufficientBalance)
	}

	if p.Id == 0 {
		return errors.New(ErrPlayerNotFound)
	}

	return err
}
