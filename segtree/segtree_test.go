package segtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SegTreeNaive struct {
	n int
	d []string
}

func NewSegTreeNaive(n int) *SegTreeNaive {
	seg := &SegTreeNaive{
		n: n,
		d: make([]string, n),
	}
	for i := 0; i < n; i++ {
		seg.d[i] = seg.e()
	}
	return seg
}

func (s *SegTreeNaive) op(a, b string) string {
	if a != "$" && b != "$" && a > b {
		panic("")
	}
	if a == "$" {
		return b
	}
	if b == "$" {
		return a
	}
	return a + b
}

func (s *SegTreeNaive) e() string {
	return "$"
}

func (s *SegTreeNaive) Set(p int, x string) {
	s.d[p] = x
}
func (s *SegTreeNaive) Get(p int) string {
	return s.d[p]
}
func (s *SegTreeNaive) Prod(l, r int) string {
	sum := s.e()
	for i := l; i < r; i++ {
		sum = s.op(sum, s.d[i])
	}
	return sum
}
func (s *SegTreeNaive) AllProd() string {
	return s.Prod(0, s.n)
}
func (s *SegTreeNaive) MaxRight(l int, f func(x string) bool) int {
	sum := s.e()
	if !f(sum) {
		panic("")
	}
	for i := l; i < s.n; i++ {
		sum = s.op(sum, s.d[i])
		if !f(sum) {
			return i
		}
	}
	return s.n
}
func (s *SegTreeNaive) MinLeft(r int, f func(x string) bool) int {
	sum := s.e()
	if !f(sum) {
		panic("")
	}
	for i := r - 1; i >= 0; i-- {
		sum = s.op(s.d[i], sum)
		if !f(sum) {
			return i + 1
		}
	}
	return 0
}

func op(a, b interface{}) interface{} {
	aa, ok := a.(string)
	if !ok {
		panic("")
	}
	bb, ok := b.(string)
	if !ok {
		panic("")
	}
	if aa != "$" && bb != "$" && aa > bb {
		panic("")
	}
	if aa == "$" {
		return bb
	}
	if bb == "$" {
		return aa
	}
	return aa + bb
}
func e() interface{} {
	return "$"
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/segtree_test.cpp#L59-L68
func TestSegtree_Zero(t *testing.T) {
	seg := New(op, e, 0)
	assert.Equal(t, "$", seg.AllProd())
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/segtree_test.cpp#L70-L97
func TestSegTree_Invalid(t *testing.T) {
	seg := New(op, e, 10)
	assert.Panics(t, func() { seg.Get(-1) })
	assert.Panics(t, func() { seg.Get(10) })

	assert.Panics(t, func() { seg.Prod(-1, -1) })
	assert.Panics(t, func() { seg.Prod(3, 2) })
	assert.Panics(t, func() { seg.Prod(0, 11) })
	assert.Panics(t, func() { seg.Prod(-1, 10) })

	assert.Panics(t, func() {
		seg.MaxRight(11, func(v interface{}) bool {
			return true
		})
	})
	assert.Panics(t, func() {
		seg.MinLeft(-1, func(v interface{}) bool {
			return true
		})
	})
	assert.Panics(t, func() {
		seg.MinLeft(0, func(v interface{}) bool {
			return false
		})
	})
}

func TestSegtree_CompareNaive(t *testing.T) {
	for n := 0; n < 30; n++ {
		seg0 := NewSegTreeNaive(n)
		seg1 := New(op, e, n)
		for i := 0; i < n; i++ {
			s := ""
			s += string(rune(i + int('a')))
			seg0.Set(i, s)
			seg1.Set(i, s)
		}
		for l := 0; l <= n; l++ {
			for r := l; r <= n; r++ {
				assert.Equal(t, seg0.Prod(l, r), seg1.Prod(l, r))
			}
		}

		for l := 0; l <= n; l++ {
			for r := l; r <= n; r++ {
				y := seg0.Prod(l, r)
				assert.Equal(t, seg0.MaxRight(l, func(x string) bool {
					return len(x) <= len(y)
				}), seg1.MaxRight(l, func(x interface{}) bool {
					xx, ok := x.(string)
					if !ok {
						panic("")
					}
					return len(xx) <= len(y)
				}))
			}
		}

		for r := 0; r <= n; r++ {
			for l := 0; l <= r; l++ {
				y := seg0.Prod(l, r)
				assert.Equal(t, seg0.MinLeft(r, func(x string) bool {
					return len(x) <= len(y)
				}), seg1.MinLeft(r, func(x interface{}) bool {
					xx, ok := x.(string)
					if !ok {
						panic("")
					}
					return len(xx) <= len(y)
				}))
			}
		}

	}
}
