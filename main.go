package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type question struct {
	quest string
	ans   []string
}

type score struct {
	correct int
	total   int
}

func main() {
	fmt.Printf("!!!!!!!!!Welcome to quiz game!!!!!!!")

	reader := bufio.NewReader(os.Stdin)

	scor := score{0, 0}

	s := []question{{
		quest: "Who was the first President of India?",
		ans:   []string{"Dr. Rajendra Prasad", "Rajendra Prasad"}},
		{
			quest: "In which year did India gain independence?",
			ans:   []string{"1947"}},
		{
			quest: "Who is known as the Father of the Nation in India?",
			ans:   []string{"Mahtma Gandhi", "Mohan Das Karamchand Gandhi"},
		},
		{
			quest: "Which ancient civilisation was located in the Indus Valley?",
			ans:   []string{"Harappan"},
		},
	}
	fmt.Println("Lets begin :)")

	for _, question := range s {
		fmt.Println(question.quest)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		correct := false
		scor.total++
		// so what should this forlop do
		// it should loop through all the possible answer and if none of them match then it should flag as incorrect
		for _, answ := range question.ans {
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
