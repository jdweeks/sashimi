package sashimi

// Board square: a8 b8 c8 d8 e8 f8 g8 h8 ... a1 b1 c1 d1 e1 f1 g1 h1
// Bit position: 63 62 61 60 59 58 57 56 ...  7  6  5  4  3  2  1  0

const (
	WHITE = iota
	BLACK
)

const (
	EMPTY         uint64 = 0
	UNIVERSE      uint64 = 0xffffffffffffffff
	FILE_A        uint64 = 0x8080808080808080
	FILE_B        uint64 = FILE_A >> 1
	FILE_C        uint64 = FILE_B >> 1
	FILE_D        uint64 = FILE_C >> 1
	FILE_E        uint64 = FILE_D >> 1
	FILE_F        uint64 = FILE_E >> 1
	FILE_G        uint64 = FILE_F >> 1
	FILE_H        uint64 = FILE_G >> 1
	RANK_1        uint64 = 0x00000000000000ff
	RANK_2        uint64 = RANK_1 << 8
	RANK_3        uint64 = RANK_2 << 8
	RANK_4        uint64 = RANK_3 << 8
	RANK_5        uint64 = RANK_4 << 8
	RANK_6        uint64 = RANK_5 << 8
	RANK_7        uint64 = RANK_6 << 8
	RANK_8        uint64 = RANK_7 << 8
	DIAG_A1_H8    uint64 = 0x0102040810204080
	DIAG_H1_A8    uint64 = 0x8040201008040201
	LIGHT_SQUARES uint64 = 0xAA55AA55AA55AA55
	DARK_SQUARES  uint64 = ^LIGHT_SQUARES
)

type Bitboard struct {
	Pawns   [2]uint64
	Knights [2]uint64
	Bishops [2]uint64
	Rooks   [2]uint64
	Queen   [2]uint64
	King    [2]uint64
}

func (board *Bitboard) InitBoard() {
	board.Pawns[WHITE] = RANK_2
	board.Pawns[BLACK] = RANK_7
	board.Knights[WHITE] = (RANK_1 & FILE_B) | (RANK_1 & FILE_G)
	board.Knights[BLACK] = (RANK_8 & FILE_B) | (RANK_8 & FILE_G)
	board.Bishops[WHITE] = (RANK_1 & FILE_C) | (RANK_1 & FILE_F)
	board.Bishops[BLACK] = (RANK_8 & FILE_C) | (RANK_8 & FILE_F)
	board.Rooks[WHITE] = (RANK_1 & FILE_A) | (RANK_1 & FILE_H)
	board.Rooks[BLACK] = (RANK_8 & FILE_A) | (RANK_8 & FILE_H)
	board.Queen[WHITE] = (RANK_1 & FILE_D)
	board.Queen[BLACK] = (RANK_8 & FILE_D)
	board.King[WHITE] = (RANK_1 & FILE_E)
	board.King[BLACK] = (RANK_8 & FILE_E)
}
