package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	filename := flag.String("csv", "problems.csv", "csv file in question, answer format")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file %s", *filename))
	}

	reader := csv.NewReader(file)

	filedata, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse given csv file")
	}

	problems := parseData(filedata)
	fmt.Println(problems)

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
			a: record[1],
		}
		problems = append(problems, problem)
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
