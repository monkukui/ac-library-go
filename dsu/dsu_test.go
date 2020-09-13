package dsu

import (
	"reflect"
	"testing"
)

func TestDisjointSetUnion(t *testing.T) {
	d := New(4)
	d.Merge(0, 1)
	if !d.Same(0, 1) {
		t.FailNow()
	}
	d.Merge(1, 2)
	if !d.Same(1, 2) {
		t.FailNow()
	}
	if !(d.Size(0) == 3) {
		t.Fatal(d.Size(0))
	}
	if d.Same(0, 3) {
		t.FailNow()
	}
	if g := d.Groups(); !reflect.DeepEqual(g, [][]int{{0, 1, 2}, {3}}) {
		t.Fatal(g)
	}
}
