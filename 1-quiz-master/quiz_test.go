package main

import (
	"strings"
	"testing"
)

func TestReadNextLine(t *testing.T) {
	t.Run("return the question", func(t *testing.T) {
		want := "5+5"
		got, err := getNextQuestion(QuizFilepath)

		if err != nil {
			t.Errorf("got error %q", err)
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("return an error if the file does not exist", func(t *testing.T) {
		_, err := getNextQuestion("a_file_that_does_not_exist")

		if err == nil {
			t.Errorf("got no error, want an error")
		}
	})
}

func TestGetUserInput(t *testing.T) {
	t.Run("return the user input", func(t *testing.T) {

		want := 10
		fakeInput := strings.NewReader("10\r\n")
		got, err := getUserAnswer(fakeInput)

		if err != nil {
			t.Errorf("got error %q", err)
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("return an error if the input is not a number", func(t *testing.T) {
		fakeInput := strings.NewReader("not_a_number\r\n")
		_, err := getUserAnswer(fakeInput)

		if err == nil {
			t.Errorf("got no error, want an error")
		}
	})
}
