package game

import "testing"

func TestRollHelperToReturnValidNumbersToBoardAndList(t *testing.T) {
	// Using the real random seed here
	usedNumbers := make([]int, 0, CellsPerSide*CellsPerSide)

	for index := 0; index < CellsPerSide*CellsPerSide; index++ {
		usedNumbers[index] = Roll(&usedNumbers)
	}

	for index := 0; index < CellsPerSide*CellsPerSide; index++ {
		if !IsIntInSlice(usedNumbers[index], &usedNumbers) {
			t.Errorf("Error rolling numbers: number in board: %d, numbers roller: %v", usedNumbers[index], usedNumbers)
			return
		}
	}
}
