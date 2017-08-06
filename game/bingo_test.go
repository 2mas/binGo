package game

import (
	"testing"
)

// Check that number of initialized boards are correct and that we get one userboard
func TestInitializingBoards(t *testing.T) {
	var wantBoards = 5
	bingo := new(Bingo)
	bingo.Initialize(wantBoards)

	if len(bingo.Boards) != wantBoards {
		t.Errorf("Unexpected number of boards in init: %d, want: %d", len(bingo.Boards), wantBoards)
		return
	}

	var foundUserBoard = false
	for _, board := range bingo.Boards {
		if board.isUserBoard && foundUserBoard {
			t.Errorf("Expecting one userboard, found multiple from init: %v", bingo.Boards)
			return
		}
		if board.isUserBoard {
			foundUserBoard = true
		}
	}

	if !foundUserBoard {
		t.Errorf("Could not find a userboard from init: %v", bingo.Boards)
		return
	}
}
