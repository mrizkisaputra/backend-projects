package helper

import (
	"fmt"
	"time"
)

var DisplayWelcome = `
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have a chances base on the difficulty level
to guess the correct number.
`

var DisplayMenu = `
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
---------------------------------------
`

func PlayerInputCmd(displayFormat string) (scan string) {
	fmt.Printf("%s: ", displayFormat)
	_, _ = fmt.Scan(&scan)
	return scan
}

func CongratulationIfCorrect(congratulation bool, totalAttempt int, computerThinkingNumber int, duration time.Duration) {
	if congratulation {
		fmt.Printf("Congratulations🎉🎉🎉\nYou guessed the correct number in %d attempts, the time you need is %v.\n\n",
			totalAttempt,
			duration)
	} else {
		fmt.Printf("Yeah you losed🥺. I'am thinking number %d.\n\n", computerThinkingNumber)
	}
}
