package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"github.com/abondar24/SocialTournamentService/blogic"
)

type Server struct {
	logic  *blogic.Logic
	router *mux.Router
}

const (
	ErrPlayerNotFound          string = "player not found"
	ErrTournamentNotFound      string = "tournament not found"
	ErrInternalError           string = "internal error"
	ErrInsufficientBalance     string = "not enough points"
	ErrBackerIsNotInTournament string = "backer is not participating in tournament"
)

func NewServer(logic *blogic.Logic) *Server {
	router := mux.NewRouter().StrictSlash(true)
	return &Server{
		logic,
		router,
	}
}

func (s *Server) RunRestServer() {

	s.router.HandleFunc("/", s.Index)
	s.router.HandleFunc("/add_player", s.AddPlayer).Methods("GET")
	s.router.HandleFunc("/take", s.Take).Methods("GET")
	s.router.HandleFunc("/fund", s.Fund).Methods("GET")
	s.router.HandleFunc("/announce_tournament",
		s.AnnounceTournament).Methods("GET")
	s.router.HandleFunc("/join_tournament", s.JoinTournament).Methods("GET")
	s.router.HandleFunc("/result_tournament", s.ResultTournament).Methods("POST")
	s.router.HandleFunc("/balance", s.Balance).Methods("GET")
	s.router.HandleFunc("/reset", s.Reset).Methods("GET")
	s.router.HandleFunc("/updatePrizes", s.UpdatePrizes).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", s.router))

}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is running")
}

func (s *Server) AddPlayer(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query()["name"][0]
	points := r.URL.Query()["points"][0]

	pts, err := strconv.Atoi(points)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.logic.AddPlayer(name, pts)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) Take(w http.ResponseWriter, r *http.Request) {
	playerId := r.URL.Query()["player_id"][0]
	points := r.URL.Query()["points"][0]

	pid, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	pts, err := strconv.Atoi(points)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.logic.Take(pid, pts)
	if err != nil {
		if err.Error() == ErrPlayerNotFound {
			log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
		}

		if err.Error() == ErrInsufficientBalance {
			log.Fatal(err)
			w.WriteHeader(http.StatusPaymentRequired)
		}

		if err.Error() == ErrInternalError {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Fund(w http.ResponseWriter, r *http.Request) {
	playerId := r.URL.Query()["player_id"][0]
	points := r.URL.Query()["points"][0]

	pid, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	pts, err := strconv.Atoi(points)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.logic.Fund(pid, pts)
	if err != nil {
		if err.Error() == ErrInternalError {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err.Error() == ErrPlayerNotFound {
			log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) AnnounceTournament(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()["name"][0]
	deposit := r.URL.Query()["deposit"][0]

	dp, err := strconv.Atoi(deposit)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.logic.AnnounceTournament(name, dp)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) JoinTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId := r.URL.Query()["tournament_id"][0]
	playerId := r.URL.Query()["player_id"][0]
	backerIds := r.URL.Query()["backerId"]

	tid, err := strconv.ParseInt(tournamentId, 10, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	pid, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	backers, err := s.convertToInt64(backerIds)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.logic.JoinTournament(tid, pid, backers)
	if err != nil {
		if err.Error() == ErrInternalError {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err.Error() == ErrTournamentNotFound {
			log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
		}

		if err.Error() == ErrPlayerNotFound {
			log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
		}

		if err.Error() == ErrInsufficientBalance {
			log.Fatal(err)
			w.WriteHeader(http.StatusPaymentRequired)
		}

		if err.Error() == ErrBackerIsNotInTournament {
			log.Fatal(err)
			w.WriteHeader(http.StatusUnauthorized)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) ResultTournament(w http.ResponseWriter, r *http.Request) {
	tournamentId := r.URL.Query()["tournament_id"][0]
	tid, err := strconv.ParseInt(tournamentId, 10, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	tr, err := s.logic.ResultTournament(tid)
	if err != nil {
		log.Fatal(err)
		if err.Error() == ErrInternalError {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err.Error() == ErrTournamentNotFound {
			log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
		}
	}

	json.NewEncoder(w).Encode(tr)
}

func (s *Server) Balance(w http.ResponseWriter, r *http.Request) {
	playerId := r.URL.Query()["player_id"][0]

	pid, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	pb, err := s.logic.Balance(pid)
	if err != nil {
		if err.Error() == ErrInternalError {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err.Error() == ErrPlayerNotFound {
			w.WriteHeader(http.StatusNotFound)
		}
	}

	json.NewEncoder(w).Encode(pb)
}

func (s *Server) Reset(w http.ResponseWriter, r *http.Request) {
	err := s.logic.Reset()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}


func (s *Server) UpdatePrizes(w http.ResponseWriter, r *http.Request) {
	tournamentId := r.URL.Query()["tournament_id"][0]
	playerId:= r.URL.Query()["player_id"][0]
	prize:= r.URL.Query()["prize"][0]

	tid, err := strconv.ParseInt(tournamentId, 10, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

    player,err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	pr,err:= strconv.Atoi(prize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.logic.UpdatePrizes(tid,player,pr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
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
