package sashimi

import (
	"math/rand"
	"time"
)

// transposition table (stores results of previous searches)
type TransTable struct {
	Keys [781]int64
}

// this table uses the Zobrist hashing technique
func NewTable() *TransTable {
	table := new(TransTable)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for j := 0; j < 781; j++ {
		table.Keys[j] = r.Int63()
	}
	return table
}
