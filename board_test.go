package sashimi

import (
	"testing"
)

func TestInit(t *testing.T) {
	board := new(Bitboard)
	board.InitBoard()
	if board.Pawns[WHITE] != RANK_2 {
		t.Error("Board initialization failed")
	}
}

func TestPrint(t *testing.T) {
	board := new(Bitboard)
	board.InitBoard()
	board.DumpBoard(board.GetAll())
}
