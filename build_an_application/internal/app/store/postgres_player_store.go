package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type PostgresPlayerStore struct {
	config *Config
	db     *sql.DB
}

func NewPostgresPlayerStore(config *Config) *PostgresPlayerStore {
	return &PostgresPlayerStore{
		config: config,
	}
}

func (s *PostgresPlayerStore) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *PostgresPlayerStore) Close() {
	s.db.Close()
}

func (s *PostgresPlayerStore) GetPlayerScore(name string) int {
	var score int
	err := s.db.QueryRow(`SELECT score FROM players_scores WHERE name=$1`, name).Scan(&score)
	if err != nil {
		log.Println(err)
	}
	return score
}

func (s *PostgresPlayerStore) RecordWin(name string) {
	_, err := s.db.Exec("UPDATE players_scores SET score = score + 1 WHERE name=$1", name)
	if err != nil {
		log.Println(err)
	}
}
