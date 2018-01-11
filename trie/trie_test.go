package trie

import (
	"testing"
	"fmt"
)

func TestBitTrie_Insert(t *testing.T) {
	bitTrie := new(BitTrie).Initialize()
	fmt.Println("after init")
	bitTrie.Insert(NewByte(7,2), "7") // 00000111
	bitTrie.Insert(NewByte(0,2), "0") // 00000000
	r := bitTrie.Search(NewByte(7,2), true)
	if r != "7" {
		t.Errorf("Expected 7 received %v", r)
	}
	r = bitTrie.Search(NewByte(15,3), true)
	if r != nil {
		t.Errorf("Expected nil received %v", r)
	}
	r = bitTrie.Search(NewByte(15,3), false)// 00001111
	if r != "7" {
		t.Errorf("Expected 7 received %v", r)
	}
}