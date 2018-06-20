package itvtree

import (
	"fmt"
	"testing"

	"github.com/silverbelial/go-trees/rbtree"
)

type Imei string

func (im Imei) Less(i rbtree.Item) bool {
	ai, ok := i.(*IntervalItem)
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

func (im Imei) Smaller(i Item) bool {
	aim, ok := i.(Imei)
	if !ok {
		//return true so that search will fail
		return true
	}
	return im < aim
}

func TestItvTree(t *testing.T) {
	itv = Init()
	itv.Insert(Imei("9000"), Imei("9005"))
	if itv.Overlapped(Imei("8000")) {
		t.Log("8000 overlaped")
		t.Fail()
	}
	if !itv.Overlapped(Imei("9002")) {
		t.Log("9002 not overlaped")
		t.Fail()
	}
}

var (
	itv *ItvTree
)

func benchmarkItvTreeInsert(b *testing.B, n int) {
	itv := Init()
	for i := 0; i < n; i++ {
		s := fmt.Sprintf("%014v", i*10000)
		e := fmt.Sprintf("%014v", i*10000+1234)
		itv.Insert(Imei(s), Imei(e))
	}
}

func BenchmarkItvTree10(b *testing.B) {
	benchmarkItvTreeInsert(b, 10)
}

func BenchmarkItvTree100(b *testing.B) {
	benchmarkItvTreeInsert(b, 100)
}

func BenchmarkItvTree1000(b *testing.B) {
	benchmarkItvTreeInsert(b, 1000)
}

func BenchmarkItvTree10000(b *testing.B) {
	benchmarkItvTreeInsert(b, 10000)
}

func BenchmarkItvTree100000(b *testing.B) {
	benchmarkItvTreeInsert(b, 100000)
}

func BenchmarkItvTreeSearch(b *testing.B) {

	itv := Init()
	for i := 0; i < 10000; i++ {
		s := fmt.Sprintf("%014v", i*10000)
		e := fmt.Sprintf("%014v", i*10000+1234)
		itv.Insert(Imei(s), Imei(e))
	}

	for i := 0; i < 500; i++ {
		iv := fmt.Sprintf("%014v", i*10000)
		itv.Overlapped(Imei(iv))
	}
}
