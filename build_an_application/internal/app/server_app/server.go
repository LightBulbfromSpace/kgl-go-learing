package server

import (
	"fmt"
	"github.com/LightBulbfromSpace/build_an_application/internal/app/store"
	"log"
	"net/http"
	"strings"
)

type PlayerServer struct {
	config *Config
	router *http.ServeMux
	store  store.PlayerStore
}

func NewServer(config *Config) *PlayerServer {
	return &PlayerServer{
		config: config,
		router: http.NewServeMux(),
	}
}

func NewServerForInMemoryStore() *PlayerServer {
	server := &PlayerServer{store: store.NewInMemoryPlayerStore(), router: http.NewServeMux()}
	server.configureRouter()
	return server
}

func (p *PlayerServer) Start() error {

	p.configureRouter()

	if err := p.configureStore(); err != nil {
		log.Println("can't configure store", err)
		return err
	}
	return nil
}

func (p *PlayerServer) configureRouter() {
	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))

}

func (p *PlayerServer) configureStore() error {
	st := store.NewPostgresPlayerStore(p.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	p.store = st
	return http.ListenAndServe(p.config.BindAddr, p.router)
}

/*func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}*/

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.proccessWin(w, player)
	}
}

func (p *PlayerServer) proccessWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
