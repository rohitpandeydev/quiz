package main

import (
	"fmt"
)

type question struct {
	quest string
	ans   string
}

func main() {
	fmt.Printf("!!!!!!!!!Welcome to quiz game!!!!!!!")
	s := make([]question, 4)
	s[0] = question{
		quest: "Who was the first President of India?",
		ans:   "Dr. Rajendra Prasad",
	}
	s[1] = question{
		quest: "In which year did India gain independence?",
		ans:   "1947",
	}
	s[2] = question{
		quest: "Who is known as the Father of the Nation in India?",
		ans:   "Mahtma Gandhi",
	}
	s[0] = question{
		quest: "Which ancient civilisation was located in the Indus Valley?",
		ans:   "Harappan",
	}
	fmt.Println("Lets begin :)")

	for _, question := range s {
		fmt.Printf(question.quest)
		var answer string
		fmt.Scanln(&answer)
		if answer == question.ans {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
	}
}
