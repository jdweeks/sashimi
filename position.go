package sashimi

type Position struct {
	Board          *Bitboard
	Side           int // WHITE or BLACK to move
	MovesSincePawn int // number of moves since last pawn push or capture
	CanCastleRight bool
	CanCastleLeft  bool
	LastMove       uint64
}

func NewPosition(b *Bitboard, s int, m int, cr bool, cl bool, lm uint64) *Position {
	pos := new(Position)
	pos.Board = b
	pos.Side = s
	pos.MovesSincePawn = m
	pos.CanCastleRight = cr
	pos.CanCastleLeft = cl
	pos.LastMove = lm
	return pos
}
