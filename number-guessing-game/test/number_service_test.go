package test

import (
	"github.com/stretchr/testify/assert"
	"number-guessing-game/cmd/services"
	"testing"
)

var service = services.NewNumberGuessingService()

func TestRandomNumber(t *testing.T) {
	for i := 1; i <= 10_000; i++ {
		randomNumber := service.RandomNumber()
		if randomNumber <= 0 || randomNumber > 100 {
			t.Errorf("Random number should be between 1 to 100. founded %d", randomNumber)
		}
	}
}

func TestPlayerGuesses(t *testing.T) {
	structs := []struct {
		Level                string
		Chances              int
		Guess                []int
		ExpectedTotalAttempt int
	}{
		{Level: "easy", Chances: 10, Guess: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, ExpectedTotalAttempt: 5},
		{Level: "medium", Chances: 5, Guess: []int{50, 25, 35, 30, 80}, ExpectedTotalAttempt: 1},
		{Level: "hard", Chances: 3, Guess: []int{25, 35, 50}, ExpectedTotalAttempt: 3},
	}

	for _, e := range structs {
		t.Run(e.Level, func(t *testing.T) {
			guesses, totalAttempt, _, err := newNumberGuessing().playerGuesses(e.Chances, e.Guess)
			assert.Nil(t, err)
			assert.True(t, guesses)
			assert.Equal(t, e.ExpectedTotalAttempt, totalAttempt)
		})
	}
}
