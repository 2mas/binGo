package main

import (
	"fmt"

	"github.com/2mas/binGo/game"
)

func main() {
	fmt.Println("Welcome to binGo, how many opponents do you want?")
	var opponents int
	_, error := fmt.Scanf("%d", &opponents)

	if error != nil {
		fmt.Println("Invalid number, defaults to zero opponents")
	}

	bingo := new(game.Bingo)
	bingo.Initialize(opponents + 1)
	bingo.RunGame()
}
