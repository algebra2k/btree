package main

import (
	"github.com/google/btree"
)

func main() {
	bt := btree.New(2)
	bt.ReplaceOrInsert(btree.Int(40))
	bt.ReplaceOrInsert(btree.Int(30))
	bt.ReplaceOrInsert(btree.Int(20))
	bt.ReplaceOrInsert(btree.Int(10))
	bt.ReplaceOrInsert(btree.Int(15))
	bt.ReplaceOrInsert(btree.Int(5)) // left leaf node full
	bt.ReplaceOrInsert(btree.Int(50))
	bt.ReplaceOrInsert(btree.Int(60))
	bt.ReplaceOrInsert(btree.Int(70))
	bt.ReplaceOrInsert(btree.Int(55))
	bt.ReplaceOrInsert(btree.Int(80)) // 串联split
	bt.ReplaceOrInsert(btree.Int(90))
	bt.ReplaceOrInsert(btree.Int(25))
	bt.ReplaceOrInsert(btree.Int(100))
}
