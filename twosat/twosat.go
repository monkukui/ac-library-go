package twosat

import internal "github.com/monkukui/ac-library-go/internal/scc"

// TwoSat defines n, answer and scc
type TwoSat struct {
	n      int
	answer []bool
	scc    *internal.SccGraph
}

// New creates a TwoSAT of n variables and 0 clauses
func New(n int) *TwoSat {
	ts := &TwoSat{
		n:      n,
		answer: make([]bool, n),
		scc:    internal.NewGraph(n * 2),
	}
	return ts
}

func (ts *TwoSat) internalJudge(f bool, a int, b int) int {
	if f {
		return a
	}
	return b
}

// AddClause adds a clause.
func (ts *TwoSat) AddClause(i int, f bool, j int, g bool) {
	if !(0 <= i && i < ts.n) {
		panic("")
	}
	if !(0 <= j && j < ts.n) {
		panic("")
	}
	ts.scc.AddEdge(2*i+ts.internalJudge(f, 0, 1), 2*j+ts.internalJudge(g, 1, 0))
	ts.scc.AddEdge(2*j+ts.internalJudge(g, 0, 1), 2*i+ts.internalJudge(f, 1, 0))
}

// Satisfiable returns True if there is a truth assignment that satisfies all clauses
func (ts *TwoSat) Satisfiable() bool {
	id := ts.scc.SccIds().Second

	for i := 0; i < ts.n; i++ {
		if id[2*i] == id[2*i+1] {
			return false
		}
		ts.answer[i] = id[2*i] < id[2*i+1]
	}
	return true
}

// Answer returns a truth assignment that satisfies all clauses of the last call of satisfiable
func (ts *TwoSat) Answer() []bool {
	return ts.answer
}
