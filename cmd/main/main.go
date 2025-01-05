package main

import (
	"log"

	"github.com/rohitpandeydev/quiz/internal/config"
	"github.com/rohitpandeydev/quiz/internal/db"
	"github.com/rohitpandeydev/quiz/internal/game"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	database, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	questions, err := database.GetQuestions()
	if err != nil {
		log.Fatalf("Failed to get questions: %v", err)
	}

	game := game.NewGame(questions)
	game.Start()
}

/*func main() {
	fmt.Printf("!!!!!!!!!Welcome to quiz game!!!!!!!\n")

	reader := bufio.NewReader(os.Stdin)

	scor := score{0, 0}

	conn, err := Connect()
	if err != nil {
		log.Fatalf("Error when connecting with db. Can't fetch the questions. Error %s", err)
	}

	QueryData(conn)

	fmt.Println("Lets begin :)")

	for _, question := range questions {
		fmt.Println(question.Question)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		correct := false
		scor.total++
		// so what should this forlop do
		// it should loop through all the possible answer and if none of them match then it should flag as incorrect
		for _, answ := range question.Answer {
			if strings.EqualFold(answer, answ) {
				correct = true
			}
		}
		if correct {
			fmt.Println("Correct")
			scor.correct++
		} else {
			fmt.Println("Incorrect")
		}
	}

	fmt.Printf("Total correct percentage %.2f", (float64(scor.correct)/float64(scor.total))*100)
}
*/
