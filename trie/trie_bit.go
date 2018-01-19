package trie

type BitTrie struct {
	root *BitTrieNode
}

type BitTrieNode struct {
	Left	*BitTrieNode
	Right	*BitTrieNode
	//Bit 	bool
	Value	interface{}
}

func(bt *BitTrie) Initialize() *BitTrie {
	bt.root = new(BitTrieNode)
	return bt
}

func(bt *BitTrie) Insert(t BitTrail, value interface{}) {
	if bt.root == nil {
		return
	}
	bt.root.Insert(t, value)
}

func(btn *BitTrieNode) Insert(t BitTrail, value interface{}) {
	if t.Empty() {
		btn.Value = value
	} else {
		bit := t.Pop()
		if bit {
			if btn.Right == nil {
				btn.Right = new(BitTrieNode)
			}
			btn.Right.Insert(t, value)
		} else {
			if btn.Left == nil {
				btn.Left = new(BitTrieNode)
			}
			btn.Left.Insert(t, value)
		}
	}
}

func(bt *BitTrie) Search(t BitTrail, full bool) interface{} {
	if bt.root == nil {
		return nil
	}
	return bt.root.Search(t, full)
}

func(btn *BitTrieNode) Search(t BitTrail, full bool) interface{} {
	if t.Empty() {
		return btn.Value
	}
	if t.Pop() {
		if btn.Right == nil {
			if full {
				return nil
			} else {
				return btn.Value
			}
		}
		return btn.Right.Search(t, full)
	}else {
		if btn.Left == nil {
			if full {
				return nil
			} else {
				return btn.Value
			}
		}
		return btn.Left.Search(t, full)
	}

}