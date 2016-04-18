package sashimi

import (
	"math/rand"
	"time"
)

// transposition table (stores results of previous searches)
type Table struct {
	Keys [781]int64
}

// this table uses the Zobrist hashing technique
func (table *Table) InitTable() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for j := 0; j < 781; j++ {
		table.Keys[j] = r.Int63()
	}
}
