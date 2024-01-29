package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Question struct that stores question with answer
type Question struct {
	question string
	answer   string
}

func main() {
	// read in quiz questions
	questions, answers := readCSV()

	defaultTimeLength := 30
	// userAnswersChan := make(chan bool, len(questions))
	// iterate through the questions in sequence
	fmt.Print(fmt.Sprintf("Sup Bitch, you have %d seconds to answer each question ", defaultTimeLength))
	score := 0
	timer := time.NewTimer(time.Duration(defaultTimeLength) * time.Second)

problemloop:
	for index, question := range questions {

		userAnswersChan := make(chan string)
		fmt.Printf(fmt.Sprintf("whats %s? :  ", question))

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			userAnswersChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Times up, moving on to the next question")
			break problemloop
		case userAnswer := <-userAnswersChan:
			if userAnswer == answers[index] {
				score = score + 1
			}
			continue
		}

	}

	fmt.Println(fmt.Sprintf("Your score was %d out of %d", score, len(questions)))
}

func readCSV() (questions []string, answers []string) {

	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	for _, record := range records {
		questions = append(questions, record[0])
		answers = append(answers, strings.TrimSpace(record[1]))
	}
	return questions, answers
}
