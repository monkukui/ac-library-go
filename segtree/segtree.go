package segtree

import (
	internal "github.com/monkukui/ac-library-go/internal/bit"
)
type SegTree struct {
	n, size, log int
	d            []interface{}
	op           func(a, b interface{}) interface{}
	e            func() interface{}
}

func (s *SegTree) update(k int) {
	s.d[k] = s.op(s.d[2*k], s.d[2*k+1])
}

func New(op func(a, b interface{}) interface{}, e func() interface{}, n int) *SegTree {
	v := make([]interface{}, n)
	for i := 0; i < n; i++ {
		v[i] = e()
	}
	return NewBySlice(op, e, v)
}
func NewBySlice(op func(a, b interface{}) interface{}, e func() interface{}, v []interface{}) *SegTree {
	n := len(v)
	log := internal.CeilPow2(n)
	size := 1 << log
	d := make([]interface{}, 2 * size)
	for i := 0; i < 2 * size; i++ {
		d[i] = e()
	}
	for i := 0; i < n; i++ {
		d[size + i] = v[i]
	}
	for i := size - 1; i >= 1; i-- {
		d[i] = op(d[2 * i], d[2 * i + 1])
	}

	return &SegTree{
		n:    n,
		size: size,
		log:  log,
		d:    d,
		op:   op,
		e:    e,
	}
}

func (s *SegTree) Set(p int, x interface{}) {
	s.d[p] = x
}

func (s *SegTree) Get(p int) interface{} {
	return s.d[p]
}

func (s *SegTree) Prod(l, r int) interface{} {
	v := s.e()
	for i := l; i < r; i++ {
		v = s.op(v, s.d[i])
	}
	return v
}

func (s *SegTree) AllProd() interface{} {
	return s.Prod(0, len(s.d))
}

func (s *SegTree) MaxRight(l int, f func(x interface{}) bool) int {
	panic("implement me!!")
}

func (s *SegTree) MinLeft(r int, f func(x interface{}) bool) int {
	panic("implement me!!")
}
