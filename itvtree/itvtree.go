package itvtree

import (
	"github.com/silverbelial/go-trees/rbtree"
)

//ItvTree implementation of interval-tree
//base on rb-tree
type ItvTree struct {
	rbTree *rbtree.RbTree
}

//IntervalItem interval-tree item
type IntervalItem struct {
	Start Item
	End   Item

	maxPoint Item
	minPoint Item
}

//Less implements rbtree.Item
func (n *IntervalItem) Less(i rbtree.Item) bool {
	an, ok := i.(*IntervalItem)
	if !ok {
		//try compare end item and targe item
		ii, ok := i.(Item)
		if !ok {
			//conversion failed , return true so that search will fail
			return true
		}
		return n.End.Smaller(ii)
	}
	// interval less: only compares start
	return n.Start.Less(an.Start)
}

//AddSuccessor implements rbtree.ExtraItem
func (n *IntervalItem) AddSuccessor(successor rbtree.ExtraItem, left bool) {
	sn, ok := successor.(*IntervalItem)
	if ok {
		if n.maxPoint.Smaller(sn.maxPoint) {
			n.maxPoint = sn.maxPoint
		}
		if sn.minPoint.Smaller(n.minPoint) {
			n.minPoint = sn.minPoint
		}
	}
}

//Recalculate implements rbtree.ExtraItem
func (n *IntervalItem) Recalculate(left, right rbtree.ExtraItem) {
	ln, ok := left.(*IntervalItem)
	if !ok {
		return
	}
	rn, ok := right.(*IntervalItem)
	if !ok {
		return
	}
	if rn.maxPoint.Smaller(ln.maxPoint) {
		n.maxPoint = ln.maxPoint
	} else {
		n.maxPoint = rn.maxPoint
	}
	n.minPoint = ln.minPoint
}

//Item comparable
type Item interface {
	Smaller(Item) bool
	Less(rbtree.Item) bool
}

//Init intialize an interval-tree
func Init() *ItvTree {
	rbt := (&rbtree.RbTree{}).Init()
	return &ItvTree{
		rbTree: rbt,
	}
}

//Insert do insertion in interval-tree
func (t *ItvTree) Insert(start, end Item) {
	// t.rbTree.

	i := &IntervalItem{Start: start, End: end, maxPoint: end, minPoint: start}
	t.rbTree.Insert(i)
}

//Overlapped check item in range
func (t *ItvTree) Overlapped(i Item) bool {
	_, found := t.rbTree.Search(i)
	return found
}

//Delete delete range from interval tree
func (t *ItvTree) Delete(start, end Item) bool {
	i := &IntervalItem{Start: start, End: end}

	_, found := t.rbTree.Delete(i)
	return found
}
