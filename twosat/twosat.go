package twosat

/*
	internal/sccに関連するコードは
	add scc #21のPull Requestから移植
	#21 がMergeされ次第internal/sccを呼ぶ形に変更
*/

//----変更予定(↓)
type sccFromToPair struct {
	first, second int
}

type sccIdPair struct {
	first  int
	second []int
}

type csr struct {
	start []int
	elist []int
}

func initCsr(n int, edges []*sccFromToPair) *csr {
	var ret csr
	ret.start = make([]int, n+1)
	ret.elist = make([]int, len(edges))
	for _, e := range edges {
		ret.start[e.first+1]++
	}
	for i := 1; i <= n; i++ {
		ret.start[i] += ret.start[i-1]
	}
	counter := make([]int, len(ret.start))
	copy(counter, ret.start)
	for _, e := range edges {
		ret.elist[counter[e.first]] = e.second
		counter[e.first]++
	}
	return &ret
}

type SccGraph struct {
	n     int
	edges []*sccFromToPair
}

func NewGraph(n int) *SccGraph {
	var s SccGraph
	s.n = n
	return &s
}

func (s *SccGraph) AddEdge(from, to int) {
	s.edges = append(s.edges, &sccFromToPair{from, to})
}

func (s *SccGraph) min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (s *SccGraph) SccIds() sccIdPair {
	g := initCsr(s.n, s.edges)
	nowOrd, groupNum := 0, 0
	visited := make([]int, 0, s.n)
	low := make([]int, s.n)
	ord := make([]int, s.n)
	ids := make([]int, s.n)
	for i := 0; i < s.n; i++ {
		ord[i] = -1
	}
	var dfs func(v int)
	dfs = func(v int) {
		low[v] = nowOrd
		ord[v] = nowOrd
		nowOrd++
		visited = append(visited, v)
		for i := g.start[v]; i < g.start[v+1]; i++ {
			to := g.elist[i]
			if ord[to] == -1 {
				dfs(to)
				low[v] = s.min(low[v], low[to])
			} else {
				low[v] = s.min(low[v], ord[to])
			}
		}
		if low[v] == ord[v] {
			for {
				u := visited[len(visited)-1]
				visited = visited[:len(visited)-1]
				ord[u] = s.n
				ids[u] = groupNum
				if u == v {
					break
				}
			}
			groupNum++
		}
	}
	for i := 0; i < s.n; i++ {
		if ord[i] == -1 {
			dfs(i)
		}
	}
	for i := 0; i < len(ids); i++ {
		ids[i] = groupNum - 1 - ids[i]
	}
	return sccIdPair{groupNum, ids}
}

//----変更予定(↑)

// TwoSat defines n, answer and scc
type TwoSat struct {
	n      int
	answer []bool
	scc    *SccGraph
}

// New creates a TwoSAT of n variables and 0 clauses
func New(n int) *TwoSat {
	ts := &TwoSat{
		n:      n,
		answer: make([]bool, n),
		scc:    NewGraph(n * 2),
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
	id := ts.scc.SccIds().second

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
