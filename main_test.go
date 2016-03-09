package jchess

import (
	"testing"
  "fmt"
)

func TestBoard(t *testing.T) {
  board := new(Bitboard)
  
  board.InitBoard()
  if board.Pawns[WHITE] != RANK_2 {
    t.Error("Board initialization failed")
  }
  
  // dump the board to stdout
  var mask uint64
  var count int = 0
  
  for mask = 1 << 63; mask > 0; mask >>= 1 {
    if count % 8 == 0 {
      fmt.Printf("\n")
    }
    count++

    if DIAG_A1_H8 & mask > 0 {
      fmt.Printf("1 ")  
    } else {
      fmt.Printf("0 ")
    }
  }
  fmt.Printf("\n\n")
}