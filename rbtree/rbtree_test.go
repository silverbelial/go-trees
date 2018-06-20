package rbtree

import (
	"testing"
)

type Imei string

func (im Imei) Less(i Item) bool {
	aim, ok := i.(Imei)
	if !ok {
		return false
	}
	return im < aim
}

func TestRbt(t *testing.T) {
	rbt := (&RbTree{}).Init()
	rbt.Insert(Imei("1234567890"))
	rbt.Insert(Imei("1234567891"))

	n, found := rbt.Search(Imei("1234567891"))
	t.Log(n)
	if !found {
		t.Fail()
	}
}
