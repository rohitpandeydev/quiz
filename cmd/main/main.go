package main

import (
	"github.com/rohitpandeydev/quiz/internal/config"
	"github.com/rohitpandeydev/quiz/internal/db"
	"github.com/rohitpandeydev/quiz/internal/game"
	"github.com/rohitpandeydev/quiz/pkg/logger"
)

func main() {
	log := logger.New(logger.INFO)

	log.Info("Starting Quiz application...")

	cfg, err := config.LoadConfig(log)
	if err != nil {
		log.Error("Failed to load config: %v", err)
		return
	}

	database, err := db.NewDB(cfg, log)
	if err != nil {
		log.Error("Failed to connect to database: %v", err)
		return
	}

	questions, err := database.GetQuestions()
	if err != nil {
		log.Error("Failed to get questions: %v", err)
		return
	}

	game := game.NewGame(questions, log)
	game.Start()
}
