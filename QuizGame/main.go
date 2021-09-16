package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const defaultProblemsFilename string = "problems.csv"

func printTask(task string) {
	fmt.Println(task)
}

func getUserAnswer() string {
	var str string

	fmt.Print("Enter your answer: ")
	fmt.Scanln(&str)

	return str
}

func printScore(problemsCount, correctAnswers int) {
	fmt.Printf("You done %d of %d\n", correctAnswers, problemsCount)
}

func playGame(records *[][]string) {
	problemsCount := len(*records)
	correctAnswers := 0

	for _, problem := range *records {
		if len(problem) < 2 {
			problemsCount--
			continue
		}

		task := problem[0]
		trueAnswer := problem[1]

		printTask(task)
		if getUserAnswer() == trueAnswer {
			correctAnswers++
		}
	}

	printScore(problemsCount, correctAnswers)
}

func main() {
	problemsFilename := defaultProblemsFilename
	if len(os.Args) >= 2 {
		problemsFilename = os.Args[1]
	}

	f, err := os.Open(problemsFilename)
	if err != nil {
		fmt.Printf("Could not open file %q, error: %s\n", problemsFilename, err.Error())
		return
	}
	defer f.Close()

	problems, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Printf("Could not read file %q, error: %s\n", problemsFilename, err.Error())
		return
	}

	playGame(&problems)
}
