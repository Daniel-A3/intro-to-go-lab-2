package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// question struct stores a single question and its corresponding answer.
type question struct {
	q, a string
}

type score int

// check handles a potential error.
// It stops execution of the program ("panics") if an error has happened.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// questions reads in questions and corresponding answers from a CSV file into a slice of question structs.
func questions() []question {
	f, err := os.Open("quiz-questions.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []question
	for _, row := range table {
		questions = append(questions, question{q: row[0], a: row[1]})
	}
	return questions
}

// ask asks a question and returns an updated score depending on the answer.
func ask(s score, questions []question, a chan score) {
	scanner := bufio.NewScanner(os.Stdin)

	for _, q := range questions {
		fmt.Println(q.q)
		fmt.Print("Enter answer: ")
		scanner.Scan()
		text := scanner.Text()
		if strings.Compare(text, q.a) == 0 {
			fmt.Println("Correct!")
			s++
		} else {
			fmt.Println("Incorrect :-(")
		}
		a <- s
	}
	close(a)

}

func main() {
	s := score(0)
	final := score(0)
	qs := questions()
	a := make(chan score, 10)
	timeout := time.After(5 * time.Second)
	go ask(s, qs, a)
	for {
		select {
		case s, ok := <-a:
			if !ok {
				fmt.Println("\nFinal score", final)
				return
			}
			final = s
		case <-timeout:
			fmt.Println("\nFinal score", final)
			return
		}

	}
}
