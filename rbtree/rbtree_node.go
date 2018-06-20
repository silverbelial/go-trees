package rbtree

const (
	cRed   = false
	cBlack = true
)

//Node tree node for Red-Black tree
type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Color  bool // use boolean for color, false for RED; true for black

	Item
}

//Item comparable
type Item interface {
	Less(Item) bool
}

//ExtraItem extra item interface when rb tree node should store extra, value which will be affected when change sub trees
//e.g. a node stores max/min value of its sub-tree
type ExtraItem interface {
	Item
	AddSuccessor(successor ExtraItem, left bool)
	Recalculate(left, right ExtraItem)
}

//returns current child is left child of its parent
func (n *Node) isLeftChild() bool {
	return n == n.Parent.Left
}

//less wrapper for node.Item.Less
func (n *Node) less(an *Node) bool {
	if n.Item == nil || an.Item == nil {
		return false
	}
	return n.Item.Less(an.Item)
}

//return node Item is ExtraItem
func (n *Node) hasExtra() bool {
	if n == nil {
		return false
	}
	_, ok := n.Item.(ExtraItem)
	return ok
}
