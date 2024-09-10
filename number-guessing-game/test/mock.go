package test

import (
	"fmt"
	"time"
)

type numberGuessingInterface interface {
	randomNumber() int
	playerGuesses(chances int, guess []int) (bool, int, time.Duration, error)
}

var instance numberGuessingInterface

type numberGuessing struct {
}

func newNumberGuessing() numberGuessingInterface {
	if instance == nil {
		instance = new(numberGuessing)
	}
	return instance
}

/* mock for testing */
func (self *numberGuessing) randomNumber() int {
	return 50 // Mocked number for consistent testing
}

/* Mock for testing */
func (self *numberGuessing) playerGuesses(chances int, guess []int) (bool, int, time.Duration, error) {
	computerThinkingNumber := self.randomNumber()
	startTime := time.Now()
	var attempt int = 1
	var success bool

	for attempt = 1; attempt <= chances && attempt <= len(guess); attempt++ {
		guessed := guess[attempt-1]
		if guessed > computerThinkingNumber {
			fmt.Printf("Incorrect! The number is less than %d\n.", guessed)
			continue
		}
		if guessed < computerThinkingNumber {
			fmt.Printf("Incorrect! The number is greater than %d\n.", guessed)
			continue
		}
		if guessed == computerThinkingNumber {
			success = true
			break
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	if !success {
		return false, attempt, duration, nil
	}
	return true, attempt, duration, nil
}
