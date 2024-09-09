package services

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"number-guessing-game/cmd/helper"
	"strconv"
	"time"
)

var instance NumberGuessingServiceInterface

type numberGuessingService struct {
}

func NewNumberGuessingService() NumberGuessingServiceInterface {
	if instance == nil {
		instance = new(numberGuessingService)
	}
	return instance
}

func (self *numberGuessingService) RandomNumber() (randomNumber int) {
	randomNumber = rand.IntN(100) + 1
	return randomNumber
}

func (self *numberGuessingService) PlayerGuesses(chances int) (bool, int, int, time.Duration, error) {
	computerThinkingNumber := self.RandomNumber()
	startTime := time.Now()
	var attempt int = 1
	var success bool

	for attempt = 1; attempt <= chances; attempt++ {
		playerInput := helper.PlayerInputCmd("Enter your guess")
		guessed, err := strconv.ParseInt(playerInput, 10, 32)
		if err != nil {
			return false, 0, 0, 0, errors.New("please enter a valid number")
		}

		if int(guessed) > computerThinkingNumber {
			fmt.Printf("Incorrect! The number is less than %d.\n\n", guessed)
			continue
		}
		if int(guessed) < computerThinkingNumber {
			fmt.Printf("Incorrect! The number is greater than %d.\n\n", guessed)
			continue
		}
		if int(guessed) == computerThinkingNumber {
			success = true
			break
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	if !success {
		return false, attempt, computerThinkingNumber, duration, nil
	}
	return true, attempt, computerThinkingNumber, duration, nil
}
