package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	filename := flag.String("csv", "problems.csv", "csv file in question, answer format")
	timelimit := flag.Int("limit", 30, "time limit in seconds")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file %s", *filename))
	}

	defer file.Close()

	reader := csv.NewReader(file)

	filedata, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse given csv file")
	}

	problems := parseData(filedata)

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	score := 0
	answerCh := make(chan string)

	for _, problem := range problems {
		fmt.Printf("%s = ", problem.q)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", score, len(problems))
			return

		case answer := <-answerCh:
			if answer == problem.a {
				fmt.Println("Correct!")
				score++
			}
		}
	}
	fmt.Printf("\nYou scored %d out of %d\n", score, len(problems))
}

type Problem struct {
	q string
	a string
}

func parseData(filedata [][]string) []Problem {
	var problems []Problem
	for _, record := range filedata {
		problem := Problem{
			q: record[0],
			a: strings.TrimSpace(record[1]),
		}
		problems = append(problems, problem)
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
