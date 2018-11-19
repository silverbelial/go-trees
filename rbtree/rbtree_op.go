package rbtree

//check the parameter node is NIL node
func (t *RbTree) isNIL(n *Node) bool {
	return t.NIL == n
}

//do left Rotation on node x
//x is the original parent node , which will be child node after the rotation
func (t *RbTree) leftRotate(x *Node) {
	// Since we are doing the left rotation, the right child should *NOT* nil.
	if t.isNIL(x.Right) {
		return
	}

	//
	// The illation of left rotation
	//
	//          |                                  |
	//          X                                  Y
	//         / \         left rotate            / \
	//        α   Y       ------------->         X   γ
	//           / \                            / \
	//          β   γ                          α   β
	//
	// It should be note that during the rotating we do not change
	// the Nodes' color.
	//
	y := x.Right
	x.Right = y.Left
	if t.isNIL(y.Left) {
		y.Left.Parent = x
	}
	y.Parent = x.Parent

	if t.isNIL(x.Parent) {
		t.root = y
	} else if x.isLeftChild() {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y

	if x.hasExtra() && y.hasExtra() {
		t.refreshExtra(x)
		t.refreshExtra(y)
	}
}

//do right Rotation on node x
//x is the original parent node , which will be child node after the rotation
func (t *RbTree) rightRotate(x *Node) {
	// Since we are doing the right rotation, the left child should *NOT* nil.
	if t.isNIL(x.Left) {
		return
	}

	//
	// The illation of right rotation
	//
	//          |                                  |
	//          X                                  Y
	//         / \         right rotate           / \
	//        Y   γ      ------------->          α   X
	//       / \                                    / \
	//      α   β                                  β   γ
	//
	// It should be note that during the rotating we do not change
	// the Nodes' color.
	//
	y := x.Left
	x.Left = y.Right
	if t.isNIL(y.Right) {
		y.Right.Parent = x
	}
	y.Parent = x.Parent

	if t.isNIL(x.Parent) {
		t.root = y
	} else if x.isLeftChild() {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Right = x
	x.Parent = y

	if x.hasExtra() && y.hasExtra() {
		t.refreshExtra(x)
		t.refreshExtra(y)
	}
}

func (t *RbTree) insert(z *Node) *Node {
	x := t.root
	y := t.NIL

	for x != t.NIL {
		y = x
		if z.less(x) { //less(z.StatusItem, x.StatusItem) {
			t.recordAddSuccessor(x, z, true)
			x = x.Left
		} else if x.less(z) { //less(x.StatusItem, z.StatusItem) {
			t.recordAddSuccessor(x, z, false)
			x = x.Right
		} else {
			//equals
			return x
		}
	}

	z.Parent = y
	if y == t.NIL {
		t.root = z
	} else if z.less(y) { //less(z.StatusItem, y.StatusItem) {
		y.Left = z
	} else {
		y.Right = z
	}

	t.count++
	t.insertFixup(z)
	return z
}

func (t *RbTree) insertFixup(z *Node) {
	for z.Parent.Color == cRed {
		//
		// Howerver, we do not need the assertion of non-nil grandparent
		// because
		//
		//  2) The root is black
		//
		// Since the color of the parent is RED, so the parent is not root
		// and the grandparent must be exist.
		//
		if z.Parent == z.Parent.Parent.Left {
			// Take y as the uncle, although it can be NIL, in that case
			// its color is BLACK
			y := z.Parent.Parent.Right
			if y.Color == cRed {
				//
				// Case 1:
				// Parent and uncle are both RED, the grandparent must be BLACK
				// due to
				//
				//  4) Both children of every red node are black
				//
				// Since the current node and its parent are all RED, we still
				// in violation of 4), So repaint both the parent and the uncle
				// to BLACK and grandparent to RED(to maintain 5)
				//
				//  5) Every simple path from root to leaves contains the same
				//     number of black nodes.
				//
				z.Parent.Color = cBlack
				y.Color = cBlack
				z.Parent.Parent.Color = cRed
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					//
					// Case 2:
					// Parent is RED and uncle is BLACK and the current node
					// is right child
					//
					// A left rotation on the parent of the current node will
					// switch the roles of each other. This still leaves us in
					// violation of 4).
					// The continuation into Case 3 will fix that.
					//
					z = z.Parent
					t.leftRotate(z)
				}
				//
				// Case 3:
				// Parent is RED and uncle is BLACK and the current node is
				// left child
				//
				// At the very beginning of Case 3, current node and parent are
				// both RED, thus we violate 4).
				// Repaint parent to BLACK will fix it, but 5) does not allow
				// this because all paths that go through the parent will get
				// 1 more black node. Then repaint grandparent to RED (as we
				// discussed before, the grandparent is BLACK) and do a right
				// rotation will fix that.
				//
				z.Parent.Color = cBlack
				z.Parent.Parent.Color = cRed
				t.rightRotate(z.Parent.Parent)
			}
		} else { // same as then clause with "right" and "left" exchanged
			y := z.Parent.Parent.Left
			if y.Color == cRed {
				z.Parent.Color = cBlack
				y.Color = cBlack
				z.Parent.Parent.Color = cRed
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					t.rightRotate(z)
				}
				z.Parent.Color = cBlack
				z.Parent.Parent.Color = cRed
				t.leftRotate(z.Parent.Parent)
			}
		}
	}
	t.root.Color = cBlack
}

// Just traverse the node from root to left recursively until left is NIL.
// The node whose left is NIL is the node with minimum value.
func (t *RbTree) min(x *Node) *Node {
	if t.isNIL(x) {
		return t.NIL
	}

	for !t.isNIL(x.Left) {
		x = x.Left
	}

	return x
}

// Just traverse the node from root to right recursively until right is NIL.
// The node whose right is NIL is the node with maximum value.
func (t *RbTree) max(x *Node) *Node {
	if t.isNIL(x) {
		return t.NIL
	}

	for t.isNIL(x.Right) {
		x = x.Right
	}

	return x
}

func (t *RbTree) search(x *Node) *Node {
	p := t.root

	for p != t.NIL {
		if p.less(x) {
			p = p.Right
		} else if x.less(p) {
			p = p.Left
		} else {
			break
		}
	}

	return p
}

//get the successor if x is deleted
//it should be noted that the x should have both left and right child when using this function
func (t *RbTree) successor(x *Node) *Node {
	if t.isNIL(x) {
		return t.NIL
	}

	// Get the minimum from the right sub-tree if it exists.
	if !t.isNIL(x.Right) {
		return t.min(x.Right)
	}

	// Get the maximum from the left sub-tree if it exists.
	if !t.isNIL(x.Left) {
		return t.max(x.Right)
	}

	return t.NIL
}

func (t *RbTree) delete(key *Node) *Node {
	//z is the target node
	z := t.search(key)

	if z == t.NIL {
		return t.NIL
	}
	ret := &Node{t.NIL, t.NIL, t.NIL, z.Color, z.Item}

	//y is the node to be removed from tree
	var y *Node
	//x is the node to check RB-tree rules
	var x *Node

	if z.Left == t.NIL || z.Right == t.NIL {
		//z has at least one child missing, delete itself; only child will replace it if exists
		y = z
	} else {
		//if both children exist, find successor to delete
		y = t.successor(z)
	}

	if y.Left != t.NIL { // left child not nil, left child will be moved up
		x = y.Left
	} else {
		// else use right child;
		// nil will be used when either child is nil
		x = y.Right
	}

	// Even if x is NIL, we do the assign. In that case all the NIL nodes will
	// change from {nil, nil, nil, BLACK, nil} to {nil, nil, ADDR, BLACK, nil},
	// but do not worry about that because it will not affect the compare
	// between Node-X with Node-NIL
	x.Parent = y.Parent

	if t.isNIL(y.Parent) { //y.Parent == t.NIL {
		t.root = x
	} else if y.isLeftChild() { //y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}

	if y != z { //successor item will be at target's place
		z.Item = y.Item
	}

	if y.Color == cBlack { //if successor is black
		t.deleteFixup(x)
	}

	t.count--
	//clean possible set
	t.NIL.Parent = t.NIL
	return ret
}

func (t *RbTree) deleteFixup(x *Node) {
	for x != t.root && x.Color == cBlack {
		if x.isLeftChild() {
			w := x.Parent.Right
			if w.Color == cRed {
				w.Color = cBlack
				x.Parent.Color = cRed
				t.leftRotate(x.Parent)
				w = x.Parent.Right
			}
			if w.Left.Color == cBlack && w.Right.Color == cBlack {
				w.Color = cRed
				x = x.Parent
			} else {
				if w.Right.Color == cBlack {
					w.Left.Color = cBlack
					w.Color = cRed
					t.rightRotate(w)
					w = x.Parent.Right
				}
				w.Color = x.Parent.Color
				x.Parent.Color = cBlack
				w.Right.Color = cBlack
				t.leftRotate(x.Parent)
				// this is to exit while loop
				x = t.root
			}
		} else { // the code below is has left and right switched from above
			w := x.Parent.Left
			if w.Color == cRed {
				w.Color = cBlack
				x.Parent.Color = cRed
				t.rightRotate(x.Parent)
				w = x.Parent.Left
			}
			if w.Left.Color == cBlack && w.Right.Color == cBlack {
				w.Color = cRed
				x = x.Parent
			} else {
				if w.Left.Color == cBlack {
					w.Right.Color = cBlack
					w.Color = cRed
					t.leftRotate(w)
					w = x.Parent.Left
				}
				w.Color = x.Parent.Color
				x.Parent.Color = cBlack
				w.Left.Color = cBlack
				t.rightRotate(x.Parent)
				x = t.root
			}
		}
	}
	x.Color = cBlack
}

//refresh extra item status
func (t *RbTree) refreshExtra(x *Node) {
	var ex, l, r ExtraItem
	var ok bool

	ex, ok = x.Item.(ExtraItem)
	if !ok {
		return
	}

	if x.Left != t.NIL {
		l, ok = x.Left.Item.(ExtraItem)
		if !ok {
			return
		}
	} else {
		l = nil
	}
	if x.Right != t.NIL {
		r, ok = x.Right.Item.(ExtraItem)
		if !ok {
			return
		}
	} else {
		r = nil
	}
	ex.Recalculate(l, r)
}

//recordAddSuccessor record add successor, change status accordingly
func (t *RbTree) recordAddSuccessor(a, s *Node, left bool) {
	if a == nil || s == nil {
		return
	}

	ea, ok := s.Item.(ExtraItem)
	if !ok {
		return
	}
	es, ok := s.Item.(ExtraItem)
	if !ok {
		return
	}

	ea.AddSuccessor(es, left)
}
