package rbtree

//RbTree represents a Red-Black tree.
type RbTree struct {
	NIL   *Node
	root  *Node
	count uint
}

//Init do tree initialization
func (t *RbTree) Init() *RbTree {
	node := &Node{nil, nil, nil, cBlack, nil}
	return &RbTree{
		NIL:   node,
		root:  node,
		count: 0,
	}
}

//Len Number of nodes in the tree.
func (t *RbTree) Len() uint { return t.count }

//Insert insert node into R-B tree
func (t *RbTree) Insert(item Item) {
	if item == nil {
		return
	}

	// Always insert a RED node
	t.insert(&Node{t.NIL, t.NIL, t.NIL, cRed, item})
}

//InsertOrGet inserts or retrieves the item in the tree. If the
//item is already in the tree then the return value will be that.
//If the item is not in the tree the return value will be the item
//you put in.
func (t *RbTree) InsertOrGet(item Item) Item {
	if item == nil {
		return nil
	}

	return t.insert(&Node{t.NIL, t.NIL, t.NIL, cRed, item}).Item
}

//Delete try delete item from R-B tree
func (t *RbTree) Delete(item Item) (Item, bool) {
	if item == nil {
		return nil, false
	}

	// The `color` field here is nobody
	n := t.delete(&Node{t.NIL, t.NIL, t.NIL, cRed, item})
	return n.Item, !t.isNIL(n)
}

//Get ???
func (t *RbTree) Get(item Item) Item {
	if item == nil {
		return nil
	}

	// The `color` field here is nobody
	ret := t.search(&Node{t.NIL, t.NIL, t.NIL, cRed, item})
	if ret == nil {
		return nil
	}

	return ret.Item
}

//Search search item node
//TODO: This is for debug, delete it in the future
func (t *RbTree) Search(item Item) (*Node, bool) {
	n := t.search(&Node{t.NIL, t.NIL, t.NIL, cRed, item})
	return n, !t.isNIL(n)
}

//Min try to find min item from R-B tree
func (t *RbTree) Min() Item {
	x := t.min(t.root)

	if x == t.NIL {
		return nil
	}

	return x.Item
}

//Max try to find max item from R-B tree
func (t *RbTree) Max() Item {
	x := t.max(t.root)

	if x == t.NIL {
		return nil
	}

	return x.Item
}
