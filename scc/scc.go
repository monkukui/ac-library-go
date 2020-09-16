package scc

import internal "github.com/monkukui/ac-library-go/internal/scc"

// SccGraph :
type SccGraph struct {
	internal *internal.SccGraph
}

// NewGraph :
func NewGraph(n int) *SccGraph {
	var ret SccGraph
	ret.internal = internal.NewGraph(n)
	return &ret
}

// AddEdge :
func (s *SccGraph) AddEdge(from, to int) {
	s.internal.AddEdge(from, to)
}

// Scc :
func (s *SccGraph) Scc() [][]int {
	return s.internal.Scc()
}
