package tests

import (
	"testing"

	"github.com/silverbelial/go-trees/itvtree"
	"github.com/silverbelial/go-trees/rbtree"
)

type Imei string

func (im Imei) Less(i rbtree.Item) bool {
	ai, ok := i.(*itvtree.IntervalItem)
	if ok {
		// used in search/overlap, compare to range start
		return im.Smaller(ai.Start)
	}
	aim, ok := i.(Imei)
	if !ok {
		//return true so that search will fail
		return true
	}
	return im < aim
}

func (im Imei) Smaller(i itvtree.Item) bool {
	aim, ok := i.(Imei)
	if !ok {
		//return true so that search will fail
		return true
	}
	return im < aim
}

func TestItvTree(t *testing.T) {
	itv := itvtree.Init()
	itv.Insert(Imei("9000"), Imei("9005"))
	itv.Insert(Imei("8005"), Imei("8010"))
	if itv.Overlapped(Imei("8000")) {
		t.Log("8000 overlapped")
	}
	if !itv.Overlapped(Imei("9002")) {
		t.Log("9002 not overlaped")
	}
}
