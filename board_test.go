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

func TestPushPawns(t *testing.T) {
	board := new(Bitboard)
	board.InitBoard()
	board.Pawns[WHITE] = PushPawns(board.Pawns[WHITE], WHITE)
	board.Pawns[BLACK] = PushPawns(board.Pawns[BLACK], BLACK)
	board.DumpBoard(board.Pawns[WHITE] | board.Pawns[BLACK])
}

func TestMoveKnights(t *testing.T) {
	board := new(Bitboard)
	board.InitBoard()
	board.Knights[WHITE] = MoveKnights(board.Knights[WHITE])
	board.Knights[BLACK] = MoveKnights(board.Knights[BLACK])
	board.DumpBoard(board.Knights[WHITE] | board.Knights[BLACK])
}
