package scc

import (
	"reflect"
	"testing"
)

func TestSimple(t *testing.T) {
	g := SccGraph(2)
	g.AddEdge(0, 1)
	g.AddEdge(1, 0)
	scc := g.Scc()
	if len(scc) != 1 {
		t.Fatal("failed Simple")
	}
}

func TestSelfLoop(t *testing.T) {
	g := SccGraph(2)
	g.AddEdge(0, 0)
	g.AddEdge(0, 0)
	g.AddEdge(1, 1)
	scc := g.Scc()
	if len(scc) != 2 {
		t.Fatal("failed Self Loop")
	}
}

func TestAlpcSample(t *testing.T) {
	g := SccGraph(6)
	edges := [][2]int{{1, 4}, {5, 2}, {3, 0}, {5, 5}, {4, 1}, {0, 3}, {4, 2}}
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}
	scc := g.Scc()
	ref := [][]int{{5}, {1, 4}, {2}, {0, 3}}
	for i, s := range scc {
		if len(s) != len(ref[i]) {
			t.Fatal("failed AlpcSample 0 ")
		}
		if !reflect.DeepEqual(scc, ref) {
			t.Fatal("failed AlpcSample 1 ")
		}
	}
}
