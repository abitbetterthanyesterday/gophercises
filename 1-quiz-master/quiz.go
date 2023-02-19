package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const QuizFilepath = "./problems.csv"

type score struct {
	Total   int
	Correct int
}

// GetNextQuestion returns the next question from the file
func getNextQuestion(filepath string) (string, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	reader := csv.NewReader(fd)
	line, err := reader.Read()
	if err != nil {
		return "", err
	}

	return readQuestionFromLine(line), nil
}

// GetUserAnswer returns the user answer
func getUserAnswer(input io.Reader) (int, error) {
	var response int

	_, err := fmt.Fscanln(input, &response)
	if err != nil {
		return -1, err
	}

	return response, nil
}

// ReadQuestionFromLine returns the question from the line
func readQuestionFromLine(line []string) string {
	// We assume the following format:
	// question,answer
	return line[0]
}

// ReadAnswerFromLine returns the answer from the line
func readAnswerFromLine(line []string) string {
	// We assume the following format:
	// question,answer
	return line[1]
}
