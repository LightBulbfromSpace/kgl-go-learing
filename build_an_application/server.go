package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const JSONContentType = "application/json"

type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
	game     Game
}

const htmlTemplatePath = "game.html"

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewServer(store PlayerStore, game Game) (*PlayerServer, error) {
	tmpl, err := template.ParseFiles(htmlTemplatePath)

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)

	}

	server := &PlayerServer{store: store, template: tmpl}
	server.configureRouter()

	server.game = game

	return server, nil
}

func (p *PlayerServer) configureRouter() {
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.playGame))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))
	p.Handler = router

}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
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

func (p *PlayerServer) playGame(w http.ResponseWriter, r *http.Request) {
	err := p.template.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem executing template %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

type PlayerServerWS struct {
	*websocket.Conn
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *PlayerServerWS {
	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Problem upgrading connection to WebSockets: %v\n", err)
	}
	return &PlayerServerWS{conn}
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	numberOfPlayersMsg := ws.WaitForMsg()
	numberOfPlayers, err := strconv.Atoi(numberOfPlayersMsg)
	if err != nil {
		log.Printf("Cannot convert number of players to int: %v\n", err)
	}

	p.game.Start(numberOfPlayers, ws)

	winner := ws.WaitForMsg()
	p.game.Finish(winner)
}

func (ws *PlayerServerWS) WaitForMsg() string {
	_, msg, err := ws.ReadMessage()
	if err != nil {
		log.Printf("Problem with reading from websocket: %v\n", err)
	}
	return string(msg)
}

func (ws *PlayerServerWS) Write(p []byte) (n int, err error) {
	err = ws.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
