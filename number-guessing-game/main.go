package main

import (
	"number-guessing-game/cmd/services"
	"number-guessing-game/cmd/views"
)

func main() {

	service := services.NewNumberGuessingService()
	view := views.NewNumberGuessingView(service)

	view.RootView()

}
