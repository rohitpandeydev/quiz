package game

import (
	"bufio"
	"os"
	"strings"

	"github.com/rohitpandeydev/quiz/internal/models"
	"github.com/rohitpandeydev/quiz/pkg/logger"
)

type Game struct {
	questions []models.Question
	score     models.Score
	reader    *bufio.Reader
	logger    *logger.Logger
}

func NewGame(questions []models.Question, log *logger.Logger) *Game {
	return &Game{
		questions: questions,
		score:     models.Score{},
		reader:    bufio.NewReader(os.Stdin),
		logger:    log,
	}
}

func (g *Game) Start() {
	g.logger.Info("Welcome to the Quiz Game!")
	g.logger.Info("------------------------")

	for i, q := range g.questions {
		g.logger.Info("\nQuestion %d: %s\n", i+1, q.Question)
		g.logger.Info("Your answer: ")

		answer, err := g.reader.ReadString('\n')
		if err != nil {
			g.logger.Error("Error reading input: %v\n", err)
			continue
		}

		// Clean the input (remove whitespace and make lowercase)
		answer = strings.TrimSpace(strings.ToLower(answer))

		// Check if answer is correct
		correct := false
		for _, validAnswer := range q.Answer {
			if strings.ToLower(validAnswer) == answer {
				correct = true
				break
			}
		}

		if correct {
			g.logger.Info("Correct!")
			g.score.Correct++
		} else {
			g.logger.Info("Wrong! The correct answer(s) were: %s\n", strings.Join(q.Answer, " or "))
		}
		g.score.Total++
	}

	// Calculate and display final score
	percentage := float64(g.score.Correct) / float64(g.score.Total) * 100
	g.logger.Info("\nQuiz completed!\n")
	g.logger.Info("You got %d out of %d questions correct (%.2f%%)\n",
		g.score.Correct, g.score.Total, percentage)
}
