package game

import (
	"fmt"
)

// Board consists of rows and cells, filled with randomizes numbers
type Board struct {
	rowCells    [CellsPerSide][CellsPerSide]int
	isUserBoard bool
	hasBingo    bool
}

// Seed creates random numbers to all cells in the board
func (b *Board) Seed() {
	usedNumbers := make([]int, 0, CellsPerSide*CellsPerSide)

	for row, cellsArray := range b.rowCells {
		// for each pair in the map, print key and value
		for cellKey := range cellsArray {
			b.rowCells[row][cellKey] = Roll(&usedNumbers)
		}
	}

	b.hasBingo = false
}

// PrintBoard prints the board and tells if it has Bingo
func (b *Board) PrintBoard() {
	fmt.Println()

	if b.isUserBoard {
		fmt.Println("Your board")
	}

	if b.hasBingo {
		if b.isUserBoard {
			fmt.Println("You got Bingo!")
		} else {
			fmt.Println("Opponent got Bingo!")
		}
	}

	for _, cellsArray := range b.rowCells {
		fmt.Println(cellsArray)
	}
}

// MarkNumberAsChecked sets the number on the board to zero
func (b *Board) MarkNumberAsChecked(number int) {
	for row, cellsArray := range b.rowCells {
		for cellKey, cellValue := range cellsArray {
			if cellValue == number {
				b.rowCells[row][cellKey] = 0
			}
		}
	}
}

// IsBingo tells if the board has bingo in horizontal or diagonal line
func (b *Board) IsBingo() bool {
	diagonalsIncline := 0
	diagonalsDecline := 0

	for rowIndex, cellsArray := range b.rowCells {
		numbersInARow := 0

		for cellIndex, cellValue := range cellsArray {
			if cellValue == 0 {
				numbersInARow++

				// Decline diagonal bingo
				if rowIndex == cellIndex {
					diagonalsDecline++
				}
				// Incline diagonal bingo
				if rowIndex+cellIndex+1 == CellsPerSide {
					diagonalsIncline++
				}

				if numbersInARow == CellsPerSide || diagonalsDecline == CellsPerSide || diagonalsIncline == CellsPerSide {
					b.hasBingo = true
				}
			}
		}
	}

	return b.hasBingo
}
