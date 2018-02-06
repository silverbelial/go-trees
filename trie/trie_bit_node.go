package trie

import "errors"

// use interface to customize implementation
type BitTrieNode interface {
	Left() (BitTrieNode, bool)
	Right() (BitTrieNode, bool)
	Value() interface{}
	GenLeft() error
	GenRight() error
	SetValue(interface{}) error
}

type BitTrieNodeMem struct {
	left	*BitTrieNodeMem
	right	*BitTrieNodeMem
	value	interface{}
}

func (btnm *BitTrieNodeMem) Left() (BitTrieNode, bool) {
	return btnm.left, btnm.left != nil
}

func (btnm *BitTrieNodeMem) Right() (BitTrieNode, bool) {
	return btnm.right, btnm.right != nil
}

func (btnm *BitTrieNodeMem) Value() interface{} {
	return btnm.value
}

func (btnm *BitTrieNodeMem) GenLeft() error {
	if btnm.left != nil {
		return errors.New("Left node already exist")
	} else {
		btnm.left = new(BitTrieNodeMem)
		return nil
	}
}

func (btnm *BitTrieNodeMem) GenRight() error {
	if btnm.right != nil {
		return errors.New("Left node already exist")
	} else {
		btnm.right = new(BitTrieNodeMem)
		return nil
	}
}

func (btnm *BitTrieNodeMem) SetValue(v interface{}) error {
	btnm.value = v
	return nil
}