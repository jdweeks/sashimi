package sashimi

import (
	"testing"
)

func TestInit(t *testing.T) {
	board := NewBoard()
	if board.Pawns[WHITE] != RANK_2 {
		t.Error("Board initialization failed")
	}
}

func TestPrint(t *testing.T) {
	board := NewBoard()
	board.PrintBoard()
	board.Knights[WHITE] = MoveKnights(board.Knights[WHITE])
	board.PrintBoard()
	board.Pawns[WHITE] = PushPawns(board.Pawns[WHITE], WHITE)
	board.PrintBoard()
}
