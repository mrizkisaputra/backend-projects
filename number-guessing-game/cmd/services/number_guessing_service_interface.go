package services

import "time"

type NumberGuessingServiceInterface interface {
	RandomNumber() int
	PlayerGuesses(chances int) (bool, int, int, time.Duration, error)
}
