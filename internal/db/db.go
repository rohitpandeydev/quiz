package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rohitpandeydev/quiz/internal/config"
	"github.com/rohitpandeydev/quiz/internal/models"
	"github.com/rohitpandeydev/quiz/pkg/logger"
)

type DB struct {
	conn   *pgx.Conn
	logger *logger.Logger
}

func NewDB(cfg *config.DBConfig, log *logger.Logger) (*DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	log.Info("Connecting to database at %s:%s", cfg.Host, cfg.Port)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &DB{conn: conn, logger: log}, nil
}

func (db *DB) GetQuestions() ([]models.Question, error) {
	db.logger.Debug("Fetching questions from database")
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
			db.logger.Error("Error scanning row: %v", err)
			return nil, err
		}
		err = json.Unmarshal(quest, &que)
		if err != nil {
			db.logger.Error("Error unmarshalling JSON: %v", err)
			return nil, err
		}
		questions = append(questions, que)
	}

	db.logger.Info("Successfully fetched %d questions", len(questions))
	return questions, nil
}
