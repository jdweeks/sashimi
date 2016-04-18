package sashimi

// push pawns one square forward
func PushPawns(pawns uint64, color int) uint64 {
	if color == WHITE {
		return pawns << 8
	} else if color == BLACK {
		return pawns >> 8
	} else {
		return pawns
	}
}

// generate knight movement patterns
func MoveKnights(knights uint64) uint64 {
	new_knights := knights
	moves := [4]uint64{15, 17, 6, 10}
	for i := 0; i < 4; i++ {
		new_knights = new_knights | (knights << moves[i]) | knights>>moves[i]
	}
	return new_knights
}
