package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/rohitpandeydev/quiz/internal/config"
	"github.com/rohitpandeydev/quiz/internal/models"
)

type DB struct {
	conn *pgx.Conn
}

func NewDB(cfg *config.DBConfig) (*DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

func (db *DB) GetQuestions() ([]models.Question, error) {
	var questions []models.Question
	rows, err := db.conn.Query(context.Background(), "SELECT id, question FROM quiz")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var quest []byte
		var que models.Question
		err := rows.Scan(&id, &quest)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(quest, &que)
		if err != nil {
			log.Fatalf("error when unmarshalling the json : %s", err)
		}
		questions = append(questions, que)
	}

	return questions, nil
}
