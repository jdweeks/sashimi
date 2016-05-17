package sashimi

type Position struct {
	Bitboard *Board
	int      Side           // WHITE or BLACK to move
	int      MovesSincePawn // number of moves since last pawn push or capture
	bool     CanCastle
	uint64   LastMove
}

func (pos *Position) SetPosition(Bitboard *b, int s, int m, bool c, unint64 lm) {
	pos.Board = b
	pos.Side = s
	pos.MovesSincePawn = m
	pos.CanCastle = c
	pos.LastMove = lm
}
