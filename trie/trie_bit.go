package trie

type BitTrie struct {
	root BitTrieNode
}

func(bt *BitTrie) Initialize(r BitTrieNode) *BitTrie {
	bt.root = r
	return bt
}

func(bt *BitTrie) Insert(t BitTrail, value interface{}) {
	if bt.root == nil {
		return
	}
	Insert(bt.root, t, value)
}

func Insert(btn BitTrieNode, t BitTrail, value interface{}) {
	if t.Empty() {
		btn.SetValue(value)
	} else {
		bit := t.Pop()
		if bit {
			right, has := btn.Right()
			if !has {
				btn.GenRight()
				right, _ = btn.Right()
			}
			Insert(right, t, value)
		} else {
			left, has := btn.Left()
			if !has {
				btn.GenLeft()
				left, _ = btn.Left()
			}
			Insert(left, t, value)
		}
	}
}

func(bt *BitTrie) Search(t BitTrail, full bool) interface{} {
	if bt.root == nil {
		return nil
	}
	return Search(bt.root, t, full)
}

func Search(btn BitTrieNode, t BitTrail, full bool) interface{} {
	if t.Empty() {
		return btn.Value()
	}
	if t.Pop() {
		right, has := btn.Right()
		if !has {
			if full {
				return nil
			} else {
				return btn.Value()
			}
		}
		return Search(right, t, full)
	}else {
		left ,has := btn.Left()
		if !has {
			if full {
				return nil
			} else {
				return btn.Value()
			}
		}
		return Search(left, t, full)
	}
}