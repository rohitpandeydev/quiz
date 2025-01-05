package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rohitpandeydev/quiz/internal/models"
)

type Game struct {
	questions []models.Question
	score     models.Score
	reader    *bufio.Reader
}

func NewGame(questions []models.Question) *Game {
	return &Game{
		questions: questions,
		score:     models.Score{},
		reader:    bufio.NewReader(os.Stdin),
	}
}

func (g *Game) Start() {
	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("------------------------")

	for i, q := range g.questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		fmt.Print("Your answer: ")

		answer, err := g.reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
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
			fmt.Println("Correct!")
			g.score.Correct++
		} else {
			fmt.Printf("Wrong! The correct answer(s) were: %s\n", strings.Join(q.Answer, " or "))
		}
		g.score.Total++
	}

	// Calculate and display final score
	percentage := float64(g.score.Correct) / float64(g.score.Total) * 100
	fmt.Printf("\nQuiz completed!\n")
	fmt.Printf("You got %d out of %d questions correct (%.2f%%)\n",
		g.score.Correct, g.score.Total, percentage)
}
