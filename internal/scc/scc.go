package internal

type sccPair struct {
	first, second int
}

type sccPair2 struct {
	first  int
	second []int
}

type csr struct {
	start []int
	elist []int
}

func initCsr(n int, edges []sccPair) *csr {
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

// SccGraph :
type SccGraph struct {
	n     int
	edges []sccPair
}

// NewSccGraph :
func NewSccGraph(n int) *SccGraph {
	var s SccGraph
	s.n = n
	return &s
}

func (s *SccGraph) numVertics() int {
	return s.n
}

// AddEdge :
func (s *SccGraph) AddEdge(from, to int) {
	s.edges = append(s.edges, sccPair{from, to})
}

func (s *SccGraph) min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (s *SccGraph) sccIds() sccPair2 {
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
	return sccPair2{groupNum, ids}
}

// Scc :
func (s *SccGraph) Scc() [][]int {
	ids := s.sccIds()
	groupNum := ids.first
	counts := make([]int, groupNum)
	for _, x := range ids.second {
		counts[x]++
	}
	groups := make([][]int, ids.first)
	for i := 0; i < groupNum; i++ {
		groups[i] = make([]int, 0, counts[i])
	}
	for i := 0; i < s.n; i++ {
		groups[ids.second[i]] = append(groups[ids.second[i]], i)
	}
	return groups
}
