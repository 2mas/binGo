package game

import (
	"fmt"
)

const (
	//CellsPerSide is the amount of rows and cells per row
	CellsPerSide = 5
)

// Bingo is the main game
type Bingo struct {
	Boards  []*Board
	IsBingo bool
}

// Initialize inits boards and seeds with numbers
func (b *Bingo) Initialize(boardCount int) {
	// Initialize boards
	for i := 0; i < boardCount; i++ {
		board := &Board{
			isUserBoard: i == 0,
		}
		board.Seed()
		b.Boards = append(b.Boards, board)
	}
}

// RunGame rolls new numbers and checks boards for bingo
func (b *Bingo) RunGame() {
	isBingo := false

	// Contains drawn numbers
	usedNumbers := make([]int, 0, CellsPerSide*CellsPerSide)

	// Draw numbers until a board gets bingo
	for isBingo == false {
		random := Roll(&usedNumbers)

		for _, board := range b.Boards {
			board.MarkNumberAsChecked(random)

			if board.IsBingo() {
				isBingo = true
			}
		}
	}

	for _, board := range b.Boards {
		board.PrintBoard()
	}

	fmt.Println()
	fmt.Println("Drawn numbers: ")
	for _, number := range usedNumbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
}
