package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"strconv"
	"github.com/abondar24/SocialTournamentService/data"
)

type Server struct {
	ds *data.MySql
	router *mux.Router

}

func NewServer(dataSource *data.MySql) *Server {
	router :=  mux.NewRouter()

	return &Server{
		dataSource,
		router,
	}
}


func (s *Server)RunRestServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", s.Index)
	//router.HandleFunc("/getPerson/{person_id}", getPerson).Methods("GET")
	//router.HandleFunc("/getPersons", getPersons).Methods("GET")
	//router.HandleFunc("/insertPerson", insertPerson).Methods("POST")
	//router.HandleFunc("/getJob/{job_id}", getJob).Methods("GET")
	//router.HandleFunc("/getJobForPerson/{person_id}", getJobForPerson).Methods("GET")
	//router.HandleFunc("/insertJob", insertJob).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is running")
}