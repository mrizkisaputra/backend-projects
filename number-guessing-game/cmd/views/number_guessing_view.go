package views

import (
	"fmt"
	"number-guessing-game/cmd/helper"
	. "number-guessing-game/cmd/services"
	"strings"
)

const (
	CHANCES_EASY_LEVEL   = 10
	CHANCES_MEDIUM_LEVEL = 5
	CHANCES_HARD_LEVEL   = 3
)

type NumberGuessingView struct {
	service NumberGuessingServiceInterface
}

func NewNumberGuessingView(numberGuessingService NumberGuessingServiceInterface) *NumberGuessingView {
	return &NumberGuessingView{
		service: numberGuessingService,
	}
}

func (self *NumberGuessingView) RootView() {
	for {
		fmt.Print(helper.DisplayWelcome)
		fmt.Print(helper.DisplayMenu)
		playerInputCmd := strings.ToLower(helper.PlayerInputCmd("Enter your choice"))

		switch playerInputCmd {
		case "1", "easy":
			{
				fmt.Print("Great! You have selected the Easy difficulty level.\nLet's start the game!\n\n")
				congratulation, totalAttempt, computerThinkingNumber, duration, err := self.service.PlayerGuesses(CHANCES_EASY_LEVEL)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				helper.CongratulationIfCorrect(congratulation, totalAttempt, computerThinkingNumber, duration)
			}
		case "2", "medium":
			{
				fmt.Print("Great! You have selected the Medium difficulty level.\nLet's start the game!\n\n")
				congratulation, totalAttempt, computerThinkingNumber, duration, err := self.service.PlayerGuesses(CHANCES_MEDIUM_LEVEL)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				helper.CongratulationIfCorrect(congratulation, totalAttempt, computerThinkingNumber, duration)
			}
		case "3", "hard":
			{
				fmt.Print("Great! You have selected the Hard difficulty level.\nLet's start the game!\n\n")
				congratulation, totalAttempt, computerThinkingNumber, duration, err := self.service.PlayerGuesses(CHANCES_HARD_LEVEL)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				helper.CongratulationIfCorrect(congratulation, totalAttempt, computerThinkingNumber, duration)
			}
		default:
			fmt.Println("Please enter a valid command!")
		}

		inputCmd := strings.ToLower(helper.PlayerInputCmd("You play again (y/n)"))
		if strings.Contains(inputCmd, "y") {
			continue
		} else {
			return
		}
	}
}
