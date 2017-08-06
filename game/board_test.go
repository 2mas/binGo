package game

import "testing"

var board *Board

func init() {
	board = &Board{
		isUserBoard: true,
	}
}

// seeds board with 1, 2, 3... to CellsPerSide*CellsPerSide (25 default)
func seedBoardIncremental(board *Board) {
	board.hasBingo = false
	i := 1

	for row, cellsArray := range board.rowCells {
		// for each pair in the map, print key and value
		for cellKey := range cellsArray {
			board.rowCells[row][cellKey] = i
			i++
		}
	}
}

func TestSeedShouldGenerateValidBoard(t *testing.T) {
	// Using the real random seed here
	board.Seed()

	usedNumbers := make([]int, 0, CellsPerSide*CellsPerSide)

	for row, cellsArray := range board.rowCells {
		// for each pair in the map, print key and value
		for cellKey := range cellsArray {
			cellNumber := board.rowCells[row][cellKey]

			if IsIntInSlice(cellNumber, &usedNumbers) {
				t.Errorf("Duplicate number in board: %d", cellNumber)
				return
			}

			if cellNumber > (CellsPerSide*CellsPerSide) || cellNumber < 1 {
				t.Errorf("Unexpected number in board: %d", cellNumber)
				return
			}

			usedNumbers = append(usedNumbers, cellNumber)
		}
	}
}

func TestIsBingo(t *testing.T) {
	seedBoardIncremental(board)

	if board.IsBingo() {
		t.Errorf("Expected IsBingo to be false")
	}
}

func TestIsBingoHorizontal(t *testing.T) {
	seedBoardIncremental(board)
	board.MarkNumberAsChecked(1)
	board.MarkNumberAsChecked(2)
	board.MarkNumberAsChecked(3)
	board.MarkNumberAsChecked(4)
	board.MarkNumberAsChecked(5)

	if !board.IsBingo() {
		t.Errorf("Expected TestIsBingoHorizontal to be true")
	}
}

func TestIsBingoDiagonalIncline(t *testing.T) {
	seedBoardIncremental(board)
	board.MarkNumberAsChecked(21)
	board.MarkNumberAsChecked(17)
	board.MarkNumberAsChecked(13)
	board.MarkNumberAsChecked(9)
	board.MarkNumberAsChecked(5)

	if !board.IsBingo() {
		t.Errorf("Expected TestIsBingoHorizontal to be true, board: %v", board)
	}
}

func TestIsBingoDiagonalDecline(t *testing.T) {
	seedBoardIncremental(board)
	board.MarkNumberAsChecked(1)
	board.MarkNumberAsChecked(7)
	board.MarkNumberAsChecked(13)
	board.MarkNumberAsChecked(19)
	board.MarkNumberAsChecked(25)

	if !board.IsBingo() {
		t.Errorf("Expected TestIsBingoHorizontal to be true, board: %v", board)
	}
}
