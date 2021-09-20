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
	d := make([]interface{}, 2*size)
	for i := 0; i < 2*size; i++ {
		d[i] = e()
	}
	for i := 0; i < n; i++ {
		d[size+i] = v[i]
	}
	for i := size - 1; i >= 1; i-- {
		d[i] = op(d[2*i], d[2*i+1])
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
	if p < 0 || s.n <= p {
		panic("")
	}
	s.d[p] = x
	p += s.size
	s.d[p] = x
	for i := 1; i <= s.log; i++ {
		s.update(p >> i)
	}
}

func (s *SegTree) Get(p int) interface{} {
	if p < 0 || s.n <= p {
		panic("")
	}
	return s.d[p+s.size]
}

func (s *SegTree) Prod(l, r int) interface{} {
	if l < 0 || r < l || s.n < r {
		panic("")
	}
	sml, smr := s.e(), s.e()
	l += s.size
	r += s.size

	for l < r {
		if l&1 > 0 {
			sml = s.op(sml, s.d[l])
			l++
		}
		if r&1 > 0 {
			r--
			smr = s.op(s.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return s.op(sml, smr)
}

func (s *SegTree) AllProd() interface{} {
	return s.d[1]
}

func (s *SegTree) MaxRight(l int, f func(x interface{}) bool) int {
	if l < 0 || s.n < l {
		panic("")
	}
	if !f(s.e()) {
		panic("")
	}
	if l == s.n {
		return s.n
	}
	l += s.size
	sm := s.e()
	for {
		for ; l%2 == 0; {
			l >>= 1
		}
		if !f(s.op(sm, s.d[l])) {
			for l < s.size {
				l = 2 * l
				if f(s.op(sm, s.d[l])) {
					sm = s.op(sm, s.d[l])
					l++
				}
			}
			return l - s.size
		}
		sm = s.op(sm, s.d[l])
		l++
		if (l & -l) != l {
			break
		}
	}
	return s.n
}

func (s *SegTree) MinLeft(r int, f func(x interface{}) bool) int {
	if r < 0 || s.n < r {
		panic("")
	}
	if !f(s.e()) {
		panic("")
	}
	if r == 0 {
		return 0
	}
	r += s.size
	sm := s.e()
	for {
		r--
		for ; r > 1 && r%2 == 1; {
			r >>= 1
		}
		if !f(s.op(s.d[r], sm)) {
			for r < s.size {
				r = 2*r + 1
				if f(s.op(s.d[r], sm)) {
					sm = s.op(s.d[r], sm)
					r--
				}
			}
			return r + 1 - s.size
		}
		sm = s.op(s.d[r], sm)
		if (r & -r) != r {
			break
		}
	}
	return 0
}
